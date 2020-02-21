package main

import (
	"fmt"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf(TerminalColors["errorColor"], err)
		}
	}()
	//reading command line args
	// commandLineArgs := os.Args[1:]
	// if len(commandLineArgs) < 1 {
	// 	fmt.Printf(TerminalColors["printColor"], "Please enter command after program name. \nE.g., "+os.Args[0]+" command_name\n")
	// 	return
	// }
	aliasFilePath := readEnvVariables()["ALIAS_FILE_PATH"].(string)
	// router(getAlias(aliasFilePath), commandLineArgs...)
	//============
	printAliases(getAlias(aliasFilePath))
	//============
}

func router(CommandsAlias map[string]string, commands ...string) {
	aliasedCommand := CommandsAlias[commands[0]]
	if len(aliasedCommand) > 0 {
		commands[0] = aliasedCommand
	}
	switch commands[0] {
	case "checkout":
		checkout(commands[1:]...)
	case "push":
		if len(commands) < 2 {
			push()
			return
		}
		execGitCmd("Couldn't execute git push", commands...)
	case "pull":
		if len(commands) < 2 {
			pull()
			return
		}
		execGitCmd("Couldn't execute git pull", commands...)
	case "commit":
		vcsFolderPath := readEnvVariables()["VSC_FOLDER_PATH"].(string)
		countCommit(vcsFolderPath, commands)
	default:
		execGitCmd("Couldn't find what you're searching for :( \n", commands...)
	}
}
