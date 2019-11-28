package main

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
)

type command struct {
	Alias   string `json:"alias"`
	Command string `json:"command"`
}

func getAlias() map[string]string {
	aliases := make(map[string]string)
	aliasFilePath, _ := filepath.Abs("../config/alias.json")
	file, err := ioutil.ReadFile(aliasFilePath)
	if err != nil {
		panic("Cannot read alias file.")
	}
	var commands []command
	err = json.Unmarshal(file, &commands)
	if err != nil {
		panic("Malformed json in alias file.")
	}
	for _, obj := range commands {
		aliases[obj.Alias] = obj.Command
	}

	return aliases
}
