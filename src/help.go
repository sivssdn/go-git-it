package main

import (
	"fmt"
	"strings"
)

//PrintAliases prints table to display all user defined aliases with their corresponding commands
func printAliases(commandsAlias map[string]string) {
	aliases := []string{}
	for alias := range commandsAlias {
		aliases = append(aliases, alias)
	}
	maxAliasLen := getMaxLength(aliases) + 2
	padding := 0
	paddingStart := strings.Join(make([]string, 4), " ")
	var aliasRow string
	fmt.Printf(TerminalColors["printColor"], " Nick Names [Alias:Command] - \n\n")
	//printing the table
	for alias := range commandsAlias {
		padding = maxAliasLen - len(alias)
		aliasRow = paddingStart + alias + strings.Join(make([]string, padding), " ") + " : " + commandsAlias[alias] + "\n"
		fmt.Printf(TerminalColors["printColor"], aliasRow)
	}
	fmt.Println()
}
