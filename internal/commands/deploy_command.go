package commands

import (
	"deploy-mez/internal/commands/helpers"
	"deploy-mez/internal/config"
	"errors"
	"fmt"
	"strings"
)

func DeployCommit(config config.Config, commit string) error {
	gitRepo := helpers.GithubRepo{
		Path: config.GitRepoPath,
	}
	// Pull repo changes
	if err := gitRepo.Pull(); err != nil {
		fmt.Println("ERROR pulling repo: ", fmt.Sprint(err))

		return err
	}
	fmt.Println("Repo has been updated successfully")

	// Get commit information
	commitDetails, err := gitRepo.GetCommitInfo(commit)
	if err != nil {
		fmt.Println("ERROR getting commit info: ", fmt.Sprint(err))

		return err
	}

	// Get all files in scope
	// For now we are not doing anything with deleted files
	files := append(commitDetails.FilesAdded, commitDetails.FilesModified...)

	// Validate XML files that are valid
	for _, file := range files {
		if strings.HasSuffix(file, ".xml") {
			filePath := helpers.GetGithubRepoFilePath(config, file)

			valid, err := helpers.IsValidXml(filePath)
			if err != nil {
				return err
			} else if !valid {
				return errors.New("Invalid XML for file: " + file)
			}
		}
	}
	fmt.Println("All XML files have been validated successfully. Valid format.")

	fmt.Println("Starting copying the files")
	if config.Backup.Enabled {
		// Implement backup functionality
	}
	for _, file := range files {
		if err := helpers.CopyFile(
			helpers.GetGithubRepoFilePath(config, file),
			helpers.GetProjectFilePath(config, file)); err != nil {
			fmt.Println("ERROR copying file: ", file)

			return err
		}

		fmt.Println("Successfully copied file: ", file)
	}

	return nil
}
