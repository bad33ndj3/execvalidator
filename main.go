package main

import (
	"log"

	"github.com/bad33ndj3/execvalidator/cmd"
	"github.com/bad33ndj3/execvalidator/pkg/execval"
)

func main() {
	// initialize the command line tool
	root := cmd.NewRootCmd()

	// add subcommands
	execValidator := execval.NewValidator()
	root.AddCommand(cmd.NewExistsCmd(execValidator))
	root.AddCommand(cmd.NewMinCmd())

	// execute the command line tool
	err := root.Execute()
	if err != nil {
		log.Println(err)
	}
}
