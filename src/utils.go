package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"
)

// TerminalColors is used for coloring terminal outputs :D
var TerminalColors = map[string]string{
	"printColor": "\033[0;36m%s\033[0m", //blue
	"errorColor": "\033[1;31m%s\033[0m", //red
}

//ExecCmd executes any terminal command
func execCmd(errMessage string, commandInput ...string) {
	command := exec.Command(commandInput[0], commandInput[1:]...)
	commandOutput, err := command.Output()
	if err != nil {
		fmt.Println("Output from terminal : ", err)
		panic(errMessage)
	}
	fmt.Printf(string(commandOutput))
}

//Executes commands withou a route in main file
func execDefaultCmd(errMessage string, commandInput ...string) {
	var commands []string
	//to account for commands aliased with multiple keywords. E.g., all => add --all
	for _, cmd := range commandInput {
		commands = append(commands, strings.Split(cmd, " ")...)
	}
	execGitCmd(errMessage, commands...)
}

//ExecCmd executes any git command
func execGitCmd(errMessage string, commandInput ...string) {
	commands := append(strings.Split("git -c color.ui=always", " "), commandInput...)
	execCmd(errMessage, commands...)
}

func getRandomString(size int) string {
	rand.Seed(time.Now().Unix())
	const letters = "abcdefghijklmnop,qrst.uvw%xyz@11"
	randString := make([]byte, size)
	for i := range randString {
		randString[i] = letters[rand.Intn(len(letters))]
	}
	return string(randString)
}

func writeStringToFile(text string, filePath string) {
	vcsfile, err := os.Create(filePath)
	defer vcsfile.Close()
	if err != nil {
		panic("Could not find file path at " + filePath)
	}
	_, errW := vcsfile.WriteString(text)
	if errW != nil {
		panic("Could not write to the specified file at " + filePath)
	}
	vcsfile.Sync()
}

func indexOf(stringSlice []string, searchText string) int {
	for key, value := range stringSlice {
		if value == searchText {
			return key
		}
	}
	return -1
}

func getMaxLength(slice []string) int {
	maxLen := 0
	for _, value := range slice {
		if len(value) > maxLen {
			maxLen = len(value)
		}
	}
	return maxLen
}
