package commands

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type GithubRepo struct {
	Path string
}

func (r *GithubRepo) Pull() {
	cmd := exec.Command(GitCommand, pullMainParameters()...)
	cmd.Dir = r.Path

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		fmt.Println(err)

		os.Exit(1)
	}
}

func (r *GithubRepo) GetCommitInfo(commit string) CommitDetails {

	return CommitDetails{
		FilesAdded:    getCommandOutputAsList(r.Path, GitCommand, getAddedFilesParameters(commit)),
		FilesModified: getCommandOutputAsList(r.Path, GitCommand, getModifiedFilesParameters(commit)),
		FilesDeleted:  getCommandOutputAsList(r.Path, GitCommand, getDeletedFilesParameters(commit)),
	}
}

func getCommandOutputAsList(projectPath string, command string, parameters []string) []string {
	cmd := exec.Command(command, parameters...)
	cmd.Dir = projectPath

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		os.Exit(1)
	}

	output := stdout.String()
	return strings.Split(strings.TrimSpace(output), "\n")
}
