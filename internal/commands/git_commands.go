package commands

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
