package main

import (
	"fmt"
	"os"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	//reading command line args
	commandLineArgs := os.Args[1:]
	if len(commandLineArgs) < 1 {
		fmt.Println(`Please enter command after [program name]. \nE.g., program_name command_name`)
		return
	}
	router(commandLineArgs...)
}

func router(commands ...string) {
	switch CommandsAlias[commands[0]] {
	case "checkout":
		if len(commands) < 2 {
			panic("Branch Identifier required for checkout.")
		}
		Checkout(commands[1])
	case "push":
		Push()
	case "pull":
		Pull()
	default:
		fmt.Println("Couldn't find what you're searching for :(")
	}
}
