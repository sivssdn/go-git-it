package main

import (
	"fmt"
)

func countCommit(vcsFolderPath string) {
	vcsFile := vcsFolderPath+"/dummy-file.txt"
	writeStringToFile(getRandomString(11), vcsFile)

	
	// execCmd(true, "Unable to execute terminal command", []string{"echo", "tmp>abc.txt"}...)
	fmt.Println("cool, will export your commit to github :)", vcsFile, getRandomString(11))
}