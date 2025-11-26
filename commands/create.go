package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/iislamgaliyev/rplate/console"
)

func RunCreateCommand() {
	createCmd := os.Args[2:]
	if len(createCmd) < 2 {
		console.PrintCreateCommandError()
		return
	}
	cmdType := createCmd[0]
	name := createCmd[1]
	inValidType := cmdType != "page" && cmdType != "feature" && cmdType != "entity"
	if inValidType {
		console.PrintCreateCommandError()
		return
	}
	invalidName := len(strings.Trim(name, " ")) == 0
	if invalidName {
		console.PrintCreateCommandError()
		return
	}
	fmt.Println("Created page | feature | entity")
}
