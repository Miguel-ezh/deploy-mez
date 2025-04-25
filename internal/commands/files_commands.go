package commands

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type FileDeployment struct {
	GitRepoPath string
	ProjectPath string
	BackupPath  string
}

func (d *FileDeployment) Deploy(commit CommitDetails) {
	for _, file := range commit.FilesAdded {
		if err := copyFile(
			buildFilePath(d.GitRepoPath, file),
			buildFilePath(d.ProjectPath, file)); err != nil {

			fmt.Println(fmt.Sprint(err) + "\n Error copying file: " + buildFilePath(d.GitRepoPath, file))
			os.Exit(1)
		}

		fmt.Println("File deployed (ADDED): " + buildFilePath(d.GitRepoPath, file))
	}

	for _, file := range commit.FilesModified {
		if err := copyFile(
			buildFilePath(d.GitRepoPath, file),
			buildFilePath(d.ProjectPath, file)); err != nil {

			fmt.Println(fmt.Sprint(err) + "\n Error copying file: " + buildFilePath(d.GitRepoPath, file))
			os.Exit(1)
		}

		fmt.Println("File deployed (MODIFIED): " + buildFilePath(d.GitRepoPath, file))
	}
}

func copyFile(source, destination string) error {
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

func buildFilePath(path, fileName string) string {
	return filepath.Join(path, fileName)
}

func (d *FileDeployment) ValidateXml(details CommitDetails) {
	for _, file := range details.FilesAdded {
		if strings.HasSuffix(file, ".xml") {
			valid, err := isValidXml(buildFilePath(d.GitRepoPath, file))
			if err != nil {
				fmt.Printf("Error validating XML: %v\n", err)
			} else if !valid {
				fmt.Println("The XML is valid.")
			}
		}
	}

	for _, file := range details.FilesModified {
		if strings.HasSuffix(file, ".xml") {
			valid, err := isValidXml(buildFilePath(d.GitRepoPath, file))
			if err != nil {
				fmt.Printf("Error validating XML: %v\n", err)
			} else if !valid {
				fmt.Println("The XML is valid.")
			}
		}
	}

}

func isValidXml(filePath string) (bool, error) {
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
