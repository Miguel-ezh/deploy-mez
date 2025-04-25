package commands

const DiffCommand = "diff"

func getDiffParameters(source, destination string) []string {
	return []string{"-q", source, destination}
}
