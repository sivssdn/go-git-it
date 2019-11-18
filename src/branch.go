package main

import (
	"fmt"
	"os/exec"
	"strings"
)

/*
	Accepts [string] substring of name of any branch
	Returns [string] first matching branch name that contains the input branch name
*/
func autoCompleteBranchName(name string) string {
	command := exec.Command("git", "branch", "-a")
	commandOutput, err := command.Output()
	if err != nil {
		panic("Couldn't execute 'git' command.")
	}
	branchesArray := strings.Split(strings.Replace(string(commandOutput), "\n", " ", -1), " ")
	for i := 1; i < len(branchesArray); i++ {
		if strings.Contains(branchesArray[i], name) {
			return branchesArray[i]
		}
	}
	return "Branch Not Found"
}

// Checkout changes current branch to the branch with matching name as the input identifier.
func Checkout(branchIdentifier string) {
	branchName := autoCompleteBranchName(branchIdentifier)
	if branchName == "Branch Not Found" {
		panic("Couldn't find any branch with given identifier")
	}
	branchName = strings.Replace(branchName, " ", "", -1)
	fmt.Println("Stashing current branch")
	execCmd("stash")
	fmt.Println("Checking out to branch : ", branchName)
	execCmd("checkout", branchName)
}

//Push updates remote with the commits on the current branch. The difference from regular push is that this function doesn't require origin to be specified.
func Push() {
	fmt.Println("Updating current branch on remote.")
	execCmd("push", "origin", "HEAD")
}

func execCmd(commandInput ...string) {
	command := exec.Command("git", commandInput...)
	commandOutput, err := command.Output()
	if err != nil {
		panic("Couldn't execute 'git " + commandInput[0] + "' command.")
	}
	fmt.Println(string(commandOutput))
}
