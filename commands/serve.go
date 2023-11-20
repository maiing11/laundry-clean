package commands

import (
	"git.enigmacamp.com/enigma-20/maher-zaenudin-mukti-umar/challenge-godb/api/routes"
	"git.enigmacamp.com/enigma-20/maher-zaenudin-mukti-umar/challenge-godb/config"
	"github.com/spf13/cobra"
)

// ServeCommand test command
type ServeCommand struct{}

func (s *ServeCommand) Short() string {
	return "serve application"
}

func (s *ServeCommand) Setup(cmd *cobra.Command) {}

func (s *ServeCommand) Run() config.CommandRunner {
	return func(
		cfg config.Config,
		router config.RequestHandler,
		route routes.Routes,
		database config.Database,
	) {
		route.Setup()
		if cfg.APIPort == "" {
			_ = router.Gin.Run()
		} else {
			_ = router.Gin.Run(":" + cfg.APIPort)
		}
	}
}

func NewServeCommand() *ServeCommand {
	return &ServeCommand{}
}
