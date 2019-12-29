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

// gets current branch name
func getCurrentBranchName() string {
	command := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	commandOutput, err := command.Output()
	if err != nil {
		panic("Cannot get current branch")
	}
	return strings.Replace(string(commandOutput), "\n", "", -1)
}

// Checkout changes current branch to the branch with matching name as the input identifier.
func checkout(branchIdentifier ...string) {
	inputBranchName := branchIdentifier[len(branchIdentifier)-1]
	branchName := autoCompleteBranchName(inputBranchName)
	branchName = strings.Replace(branchName, " ", "", -1)
	fmt.Printf(TerminalColors["printColor"], "Stashing current branch\n")
	execGitCmd("Couldn't execute git stash", "stash")
	fmt.Printf(TerminalColors["printColor"], "Checking out to branch : "+branchName+"\n")
	branchIdentifier[len(branchIdentifier)-1] = branchName
	checkoutCommand := []string{"checkout"}
	checkoutCommand = append(checkoutCommand, branchIdentifier...)
	execGitCmd("Couldn't execute git checkout on the mentioned branch", checkoutCommand...)
}

//Push updates remote with the commits on the current branch. The difference from regular push is that this function doesn't require origin to be specified.
func push() {
	fmt.Printf(TerminalColors["printColor"], "Updating current branch on remote.\n")
	execGitCmd("Couldn't execute git push. This command", "push", "origin", getCurrentBranchName())
}

//Pull updates local branch with the contents of the remote branch without specifying origin or upstream.
func pull() {
	fmt.Printf(TerminalColors["printColor"], "Stashing current branch\n")
	execGitCmd("Couldn't execute git stash", "stash")
	fmt.Printf("Pulling from remote branch.\n")
	execGitCmd("Couldn't execute git pull.", "pull", "origin", getCurrentBranchName())
}
