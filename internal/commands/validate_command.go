package commands

import (
	"deploy-mez/internal/commands/helpers"
	"deploy-mez/internal/config"
	"fmt"
)

func ValidateCommit(config config.Config, commit string) error {
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
	files := append(commitDetails.FilesAdded, commitDetails.FilesModified...)
	for _, file := range files {
		_, err := helpers.IsSameFile(helpers.GetGithubRepoFilePath(config, file),
			helpers.GetProjectFilePath(config, file))
		if err != nil {

			fmt.Println("ERROR validating file: ", file)
			return err
		}
	}

	fmt.Println("Successfully validated commit: ", commit)
	return nil
}
