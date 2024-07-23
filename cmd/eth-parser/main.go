package main

import (
	"os"

	"github.com/MonikaCat/eth-parser/cmd"
	"github.com/MonikaCat/eth-parser/config"
)

func main() {
	// initialise config
	config := cmd.NewConfig("eth-parser").
		WithConfig(&config.Config{})

	// build the default runner using provided
	// configuration and run the commands
	runner := cmd.Runner(config)
	err := runner.Execute()
	if err != nil {
		// exit the program if execution fails
		os.Exit(1)
	}
}
