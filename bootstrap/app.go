package bootstrap

import (
	"git.enigmacamp.com/enigma-20/maher-zaenudin-mukti-umar/enigma-laundry-clean/commands"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "laundry-clean",
	Short: "Clean architecture using gin framework",
	Long: `

╭━━━╮╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╭╮╱╱╱╱╱╱╱╱╱╱╱╱╱╭╮╱╱╱╱╱╭━━━┳╮
┃╭━━╯╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱┃┃╱╱╱╱╱╱╱╱╱╱╱╱╱┃┃╱╱╱╱╱┃╭━╮┃┃
┃╰━━┳━╮╭┳━━┳╮╭┳━━╮╱╱┃┃╱╱╭━━┳╮╭┳━╮╭━╯┣━┳╮╱╭┫┃╱╰┫┃╭━━┳━━┳━╮
┃╭━━┫╭╮╋┫╭╮┃╰╯┃╭╮┣━━┫┃╱╭┫╭╮┃┃┃┃╭╮┫╭╮┃╭┫┃╱┃┃┃╱╭┫┃┃┃━┫╭╮┃╭╮╮
┃╰━━┫┃┃┃┃╰╯┃┃┃┃╭╮┣━━┫╰━╯┃╭╮┃╰╯┃┃┃┃╰╯┃┃┃╰━╯┃╰━╯┃╰┫┃━┫╭╮┃┃┃┃
╰━━━┻╯╰┻┻━╮┣┻┻┻╯╰╯╱╱╰━━━┻╯╰┻━━┻╯╰┻━━┻╯╰━╮╭┻━━━┻━┻━━┻╯╰┻╯╰╯
╱╱╱╱╱╱╱╱╭━╯┃╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╭━╯┃
╱╱╱╱╱╱╱╱╰━━╯╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╰━━╯ 
                                         		
This is a command runner or cli for api architecture in golang. 
Using this we can use underlying dependency injection container for running scripts. 
Main advantage is that, we can use same usecases, repositories, infrastructure present in the application itself`,
	TraverseChildren: true,
}

// App root of application
type App struct {
	*cobra.Command
}

func NewApp() App {
	cmd := App{
		Command: rootCmd,
	}
	cmd.AddCommand(commands.GetSubCommands(CommonModules)...)
	return cmd
}

var RootApp = NewApp()
