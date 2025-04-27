package helpers

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"strings"
)

type CommitDetails struct {
	FilesAdded    []string
	FilesModified []string
	FilesDeleted  []string
}

const GitCommand string = "git"

func pullMainParameters() []string {
	return []string{"pull", "origin", "main"}
}

func getAddedFilesParameters(commit string) []string {
	return []string{"show", "--name-only", "--pretty=format:", "--diff-filter=A", commit}
}

func getModifiedFilesParameters(commit string) []string {
	return []string{"show", "--name-only", "--pretty=format:", "--diff-filter=M", commit}
}

func getDeletedFilesParameters(commit string) []string {
	return []string{"show", "--name-only", "--pretty=format:", "--diff-filter=D", commit}
}

type GithubRepo struct {
	Path string
}

func (r *GithubRepo) Pull() error {
	cmd := exec.Command(GitCommand, pullMainParameters()...)
	cmd.Dir = r.Path

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func (r *GithubRepo) GetCommitInfo(commit string) (CommitDetails, error) {
	filesAdded, err := getCommandOutputAsList(r.Path, GitCommand, getAddedFilesParameters(commit))
	if err != nil {
		return CommitDetails{}, err
	}
	filesModified, err := getCommandOutputAsList(r.Path, GitCommand, getModifiedFilesParameters(commit))
	if err != nil {
		return CommitDetails{}, err
	}
	filesDeleted, err := getCommandOutputAsList(r.Path, GitCommand, getDeletedFilesParameters(commit))
	if err != nil {
		return CommitDetails{}, err
	}

	return CommitDetails{
		FilesAdded:    filesAdded,
		FilesModified: filesModified,
		FilesDeleted:  filesDeleted,
	}, nil
}

func getCommandOutputAsList(projectPath string, command string, parameters []string) ([]string, error) {
	cmd := exec.Command(command, parameters...)
	cmd.Dir = projectPath

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		errMessage := stderr.String()

		return nil, errors.New(strings.TrimSpace(errMessage))
	}

	output := stdout.String()
	return strings.Split(strings.TrimSpace(output), "\n"), nil
}
