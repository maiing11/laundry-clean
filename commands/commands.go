package commands

import (
	"context"
	"log"

	"git.enigmacamp.com/enigma-20/maher-zaenudin-mukti-umar/challenge-godb/config"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var cmds = map[string]config.Command{
	"app:serve": NewServeCommand(),
}

// GetSubCommands gives a list of sub commands
func GetSubCommands(opt fx.Option) []*cobra.Command {
	var subCommands []*cobra.Command
	for name, cmd := range cmds {
		subCommands = append(subCommands, WrapSubCommand(name, cmd, opt))
	}
	return subCommands
}

func WrapSubCommand(name string, cmd config.Command, opt fx.Option) *cobra.Command {
	wrappedCmd := &cobra.Command{
		Use:   name,
		Short: cmd.Short(),
		Run: func(c *cobra.Command, args []string) {
			opts := fx.Options(
				fx.Invoke(cmd.Run()),
			)
			ctx := context.Background()
			app := fx.New(opt, opts)
			err := app.Start(ctx)
			defer app.Stop(ctx)
			if err != nil {
				log.Fatal(err)
			}
		},
	}
	cmd.Setup(wrappedCmd)
	return wrappedCmd
}
