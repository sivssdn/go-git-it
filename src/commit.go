package main

import (
	"os"
	"strings"
)

//commits a change to the vcs folder of git and then proceeds to normal commit behviour in current repo
func countCommit(vcsFolderPath string, commitCommand []string) {
	//changing contents of file at git repo
	vcsFileName := "dummy-file.txt"
	vcsFilePath := vcsFolderPath + string(os.PathSeparator) + vcsFileName
	writeStringToFile(getRandomString(11), vcsFilePath)

	//commiting changes to git repo
	command := "--git-dir "+vcsFolderPath + string(os.PathSeparator) + ".git/ --work-tree "+vcsFolderPath+" add --all"
	execGitCmd("Unable to add files at git folder", strings.Split(command, " ")...)
	commitMessageIndex := indexOf(commitCommand, "-m")
	commitMessage := "No Commit Message"
	if (commitMessageIndex != -1 && commitMessageIndex+1 < len(commitCommand)){
		commitMessage = commitCommand[indexOf(commitCommand, "-m")+1]
	}
	command = "--git-dir "+vcsFolderPath + string(os.PathSeparator) + ".git/ --work-tree " + vcsFolderPath + " commit -m " + strings.Replace(commitMessage, " ", "-", -1)
	execGitCmd("Unable to commit changes at git folder", strings.Split(command, " ")...)

	//carry on what the commit on current directory
	execGitCmd("Unable to commit changes at current folder", commitCommand...)
}