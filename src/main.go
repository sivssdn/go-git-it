package main

import (
	"fmt"
	"os"
)

func main() {
	commandLineArgs := os.Args[1:]
	if len(commandLineArgs) < 1 {
		fmt.Println(`Please enter command after [program name]. \nE.g., program_name command_name`)
		return
	}
	fmt.Println(commandLineArgs)
}
