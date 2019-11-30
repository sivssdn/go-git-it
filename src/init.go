package main

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"runtime"
)

type command struct {
	Alias   string `json:"alias"`
	Command string `json:"command"`
}

func getAlias() map[string]string {
	aliases := make(map[string]string)
	_, filename, _, status := runtime.Caller(0)
	if !status {
		panic("Cannot read alias file.")
	}
	aliasFilePath, _ := filepath.Abs(filepath.Dir(filepath.Dir(filename)) + "/config/alias.json")
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
