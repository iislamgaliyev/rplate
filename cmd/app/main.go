package main

import (
	"os"

	"github.com/iislamgaliyev/rplate/commands"
	"github.com/iislamgaliyev/rplate/console"
)

func main() {

	if len(os.Args) < 2 {
		commands.RunHelpCommand()
		return
	}
	mainCmd := os.Args[1]

	switch mainCmd {
	case "-h":
		commands.RunHelpCommand()
	case "-gen":
		commands.RunGenerateTemplateCommand()
	case "-c":
		commands.RunCreateCommand()
	default:
		console.PrintAllCommandsDescription()
	}

}
