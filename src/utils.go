package main

import(
	"fmt"
	"os/exec"
	"strings"
	"math/rand"
)

// TerminalColors is used for coloring terminal outputs :D
var TerminalColors = map[string]string{
	"printColor": "\033[0;36m%s\033[0m", //blue
	"errorColor": "\033[1;31m%s\033[0m", //red
}

//ExecCmd executes any terminal command
func execCmd(errMessage string, commandInput ...string) {
	commands := append(strings.Split("-c color.ui=always", " "), commandInput...)
	command := exec.Command("git", commands...)
	commandOutput, err := command.Output()
	if err != nil {
		fmt.Println("Output from git : ", err)
		panic(errMessage)
	}
	fmt.Printf(string(commandOutput))
}

func getRandomString(size int) string {
	const letters = "abcdefghijklmnop,qrst.uvw%xyz@"
	randString := make([]byte, size)
	for i := range randString{
		randString[i] = letters[rand.Intn(len(letters))]
	}
	return string(randString)
}