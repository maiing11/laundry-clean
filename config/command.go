package config

import "github.com/spf13/cobra"

type CommandRunner interface{}

// Command interface berguna untuk implementasi sub-commands di system.
type Command interface {
	// Short mengembalikan string tentang deskripsi pendek dari command.
	// string ini akan tampil di help screen dari cobra command.
	Short() string

	// setup is used to setup flags or pre-run steps for the command.
	//
	// for example,
	// cmd.Flag().IntVarP("flagname", "f", 1234, "help message")

	Setup(cmd *cobra.Command)

	// Run runs the command runner
	// run returns command runner which is a function with dependency
	// injected arguments.
	//
	// For example,
	//  Command{
	//   Run: func(l config.Logger) {
	// 	   l.Info("yay it's work")
	// 	 },
	//  }
	//
	Run() CommandRunner
}
