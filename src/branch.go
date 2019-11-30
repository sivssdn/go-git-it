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
	return name
}

// Checkout changes current branch to the branch with matching name as the input identifier.
func checkout(branchIdentifier ...string) {
	inputBranchName := branchIdentifier[len(branchIdentifier)-1]
	branchName := autoCompleteBranchName(inputBranchName)
	branchName = strings.Replace(branchName, " ", "", -1)
	fmt.Printf(TerminalColors["printColor"], "Stashing current branch\n")
	execCmd("Couldn't execute git stash", "stash")
	fmt.Printf(TerminalColors["printColor"], "Checking out to branch : "+branchName+"\n")
	branchIdentifier[len(branchIdentifier)-1] = branchName
	checkoutCommand := []string{"checkout"}
	checkoutCommand = append(checkoutCommand, branchIdentifier...)
	execCmd("Couldn't execute git checkout on the mentioned branch", checkoutCommand...)
}

//Push updates remote with the commits on the current branch. The difference from regular push is that this function doesn't require origin to be specified.
func push() {
	fmt.Printf(TerminalColors["printColor"], "Updating current branch on remote.\n")
	execCmd("Couldn't execute git push", "push", "origin", "HEAD")
}

//Pull updates local branch with the contents of the remote branch without specifying origin or upstream.
func pull() {
	fmt.Printf("Pulling from remote branch.\n")
	execCmd("Couldn't execute git pull, please check if changes are stashed on current branch", "pull", "origin", "HEAD")
}

//ExecCmd executes any git command
func execCmd(errMessage string, commandInput ...string) {
	commands := append(strings.Split("-c color.ui=always", " "), commandInput...)
	command := exec.Command("git", commands...)
	commandOutput, err := command.Output()
	if err != nil {
		panic(errMessage)
	}
	fmt.Printf(string(commandOutput))
}
