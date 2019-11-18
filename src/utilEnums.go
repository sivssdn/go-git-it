package main

// CommandsAlias is used as dictionary for customizing git commands
var CommandsAlias = map[string]string{
	"goon": "checkout",
	"pull": "pull",
	"push": "push",
}

// TerminalColors is used for coloring terminal outputs :D
var TerminalColors = map[string]string{
	"printColor": "\033[0;36m%s\033[0m", //blue
	"errorColor": "\033[1;31m%s\033[0m", //red
}
