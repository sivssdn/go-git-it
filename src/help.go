package main

import (
	"reflect"
)

//PrintAliases prints table to display all user defined aliases with their corresponding commands
func printAliases(commandsAlias map[string]string) {
	aliases := reflect.ValueOf(commandsAlias).MapKeys()
	maxAliasLen, maxCommandLen := 0, 0
	for _, value := range commandsAlias {
		if len(value) > maxCommandLen {
			maxCommandLen = len(value)
		}
	}
}
