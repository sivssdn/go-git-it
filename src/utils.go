package main

import(
	"fmt"
	"os/exec"
	"strings"
	"math/rand"
	"os"
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

//ExecCmd executes any git command
func execGitCmd(errMessage string, commandInput ...string){
	commands := append(strings.Split("git -c color.ui=always", " "), commandInput...)
	execCmd(errMessage, commands...)
}

func getRandomString(size int) string {
	rand.Seed(time.Now().Unix())
	const letters = "abcdefghijklmnop,qrst.uvw%xyz@"
	randString := make([]byte, size)
	for i := range randString{
		randString[i] = letters[rand.Intn(len(letters))]
	}
	return string(randString)
}

func writeStringToFile(text string, filePath string){
	vcsfile, err := os.Create(filePath)
	defer vcsfile.Close()
	if err != nil {
		panic("Could not find file path at "+filePath)
	}
	_, errW := vcsfile.WriteString(text)
	if errW != nil {
		panic("Could not write to the specified file at "+filePath)
	}
	vcsfile.Sync()
}