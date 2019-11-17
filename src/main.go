package main

import (
	"fmt"
	"os"
)

func main() {
	defer func(){
		if err := recover(); err != nil {
			fmt.Println("Something unexpected happened \n", err)
		}
	}()
	
	//reading command line args
	commandLineArgs := os.Args[1:]
	if len(commandLineArgs) < 1 {
		fmt.Println(`Please enter command after [program name]. \nE.g., program_name command_name`)
		return
	}
	fmt.Println(commandLineArgs)
	AutoComplete("Hi branches!")
	
}
