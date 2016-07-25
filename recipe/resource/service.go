package resource

import (
	"github.com/k0kubun/itamae-go/logger"
)

type Service struct {
	Base
	Name     string
	Provider string
}

func (s *Service) Apply() {
	logger.Debug("file[" + s.Name + "] will not change")
}

func (s *Service) DryRun() {
	logger.Color(logger.Green, func() {
		logger.Info(s.Resource + " executed will change from 'false' to 'true'")
	})
}
