package main

import (
	"fmt"
	"os"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf(TerminalColors["errorColor"], err)
		}
	}()

	//reading command line args
	commandLineArgs := os.Args[1:]
	if len(commandLineArgs) < 1 {
		fmt.Printf(TerminalColors["printColor"], `Please enter command after [program name]. \nE.g., program_name command_name\n`)
		return
	}
	router(commandLineArgs...)
}

func router(commands ...string) {
	switch CommandsAlias[commands[0]] {
	case "checkout":
		Checkout(commands[1:]...)
	case "push":
		if len(commands) < 2 {
			Push()
			return
		}
		commands[0] = "push"
		ExecCmd("Couldn't execute git push", commands...)
	case "pull":
		if len(commands) < 2 {
			Pull()
			return
		}
		commands[0] = "pull"
		ExecCmd("Couldn't execute git pull", commands...)
	default:
		ExecCmd("Couldn't find what you're searching for :(", commands...)
	}
}
