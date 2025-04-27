package helpers

import (
	"deploy-mez/internal/config"
	"encoding/xml"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func IsValidXml(filePath string) (bool, error) {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return false, err
	}
	defer file.Close()

	// Create an XML decoder
	decoder := xml.NewDecoder(file)

	// Attempt to decode the XML
	for {
		_, err := decoder.Token()
		if err != nil {
			if err.Error() == "EOF" {
				// Reached the end of the file, XML is valid
				return true, nil
			}
			// Return false if an error occurs
			return false, err
		}
	}
}

func GetGithubRepoFilePath(config config.Config, file string) string {
	return filepath.Join(config.GitRepoPath, file)
}

func GetProjectFilePath(config config.Config, file string) string {
	if config.Project.NormalizedPath != "" {
		normalizedPath := strings.Replace(file, config.Project.NormalizedPath, "", 1)

		return filepath.Join(config.Project.ProjectPath, normalizedPath)
	}

	return filepath.Join(config.Project.ProjectPath, file)
}

func GetBackupFilePath(config config.Config, file string) string {
	return filepath.Join(config.Backup.BackupPath, file)
}

func CopyFile(source, destination string) error {
	destinationDir := filepath.Dir(destination)
	if err := os.MkdirAll(destinationDir, os.ModePerm); err != nil {
		return err
	}

	srcFile, err := os.Open(source)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	destFile, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	return err
}
