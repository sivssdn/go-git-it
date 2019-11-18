package main

import (
	"os/exec"
	"strings"
)

/*
	Accepts [string] substring of name of any branch
	Returns [string] first matching branch name that contains the input branch name
*/
func AutoCompleteBranchName(name string) string {
	command := exec.Command("git", "branch", "-a")
	commandOutput, err := command.Output()
	if err != nil {
		panic(err)
	}
	branchesArray := strings.Split(strings.Replace(string(commandOutput), "\n", " ", -1), " ")
	for i := 1; i < len(branchesArray); i++ {
		if strings.Contains(branchesArray[i], name) {
			return branchesArray[i]
		}
	}
	return "Branch Not Found"
}
