package helpers

import (
	"bytes"
	"errors"
	"os/exec"
	"strings"
)

const DiffCommand = "diff"

func getDiffParameters(source, destination string) []string {
	return []string{"-q", source, destination}
}

func IsSameFile(source, destination string) (bool, error) {
	cmd := exec.Command(DiffCommand, getDiffParameters(source, destination)...)

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		outputMessage := strings.TrimSpace(stdout.String())

		return false, errors.New(strings.TrimSpace(outputMessage))
	}

	return true, nil
}
