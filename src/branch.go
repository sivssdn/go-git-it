package main

import (
	"fmt"
	"os/exec"
)

func AutoComplete(name string){
	fmt.Println(name)
	command := exec.Command("git", "branch")
	branches, err := command.Output()

	if err != nil {
		panic(err)
	}
	fmt.Println(string(branches))
}