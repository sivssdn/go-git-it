package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/joho/godotenv"
)

type command struct {
	Alias   string `json:"alias"`
	Command string `json:"command"`
}

//gets alias and their respective commands from (alias.json) file.
func getAlias(aliasFilePath string) map[string]string {
	aliases := make(map[string]string)
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

//read variables from .env file in any format
func readEnvVariables() map[string]interface{} {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Unable to read .env file.")
	}
	return map[string]interface{} {
		"ALIAS_FILE_PATH": os.Getenv("ALIAS_FILE_PATH"),
		"VSC_FOLDER_PATH": os.Getenv("VSC_FOLDER_PATH"),
	}
}
