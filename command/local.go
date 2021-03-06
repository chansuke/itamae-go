package command

import (
	"flag"
	"log"
	"strings"

	"github.com/k0kubun/itamae-go/itamae"
	"github.com/k0kubun/itamae-go/logger"
	"github.com/k0kubun/itamae-go/recipe"
)

type LocalCommand struct {
	Meta

	dryRun   bool
	nodeJson string
	logLevel string
	recipes  []string
}

func (c *LocalCommand) Run(args []string) int {
	err := c.parseArgs(args)
	if err != nil {
		return 1
	}
	logger.Info("Starting itamae...")
	logger.SetLogLevel(c.logLevel)

	context := recipe.NewContext()
	defer context.Close()

	context.LoadJson(c.nodeJson)
	for _, file := range c.recipes {
		context.LoadRecipe(file)
	}

	if c.dryRun {
		itamae.DryRun(context.Resources())
	} else {
		itamae.Apply(context.Resources())
	}
	return 0
}

func (c *LocalCommand) parseArgs(args []string) error {
	flags := flag.NewFlagSet("itamae-go", flag.ContinueOnError)

	flags.BoolVar(&c.dryRun, "n", false, "Dry run")
	flags.BoolVar(&c.dryRun, "dry-run", false, "Dry run")
	flags.StringVar(&c.nodeJson, "j", "", "Node JSON")
	flags.StringVar(&c.nodeJson, "node-json", "", "Node JSON")
	flags.StringVar(&c.logLevel, "l", "", "Log level")
	flags.StringVar(&c.logLevel, "log-level", "", "Log level")

	if err := flags.Parse(args); err != nil {
		return err
	}
	for 0 < flags.NArg() {
		c.recipes = append(c.recipes, flags.Arg(0))
		flags.Parse(flags.Args()[1:])
	}
	if len(c.recipes) == 0 {
		log.Fatal("Please specify recipe files.")
	}
	return nil
}

func (c *LocalCommand) Synopsis() string {
	return "Run itamae locally"
}

func (c *LocalCommand) Help() string {
	helpText := `
Usage:
  itamae-go local RECIPE [RECIPE...]

Options:
  -j, [--node-json=NODE_JSON]
  -n, [--dry-run]
  -l, [--log-level=LOG_LEVEL]

Run Itamae locally
`
	return strings.TrimSpace(helpText)
}
