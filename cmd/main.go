package main

import (
	"deploy-mez/internal/commands"
	"deploy-mez/internal/config"
	"flag"
	"fmt"
	"os"
)

func main() {
	// Define command-line flags
	deploy := flag.String("d", "", "Deploy files from commit")
	validate := flag.String("v", "", "Validate files from commit")
	help := flag.Bool("help", false, "Display help information")

	// Parse the flags
	flag.Parse()

	// Display help information if -help is provided
	if *help {
		fmt.Println("Deploy")
		fmt.Println("Usage: program [-d <commit_hash> | -v <commit_hash> | -help]")
		fmt.Println("Options:")
		fmt.Println("  -d <commit_hash>    Process configuration text")
		fmt.Println("  -v <commit_hash>    Process validation text")
		fmt.Println("  -help               Display this help message")
		fmt.Println("Note: -d and -v cannot be used together")
		os.Exit(0)
	}

	// Check if both flags are provided
	if *deploy != "" && *validate != "" {
		fmt.Println("Error: Cannot use -d and -v together")
		fmt.Println("Use -help for usage information")
		os.Exit(1)
	}

	// Check if neither flag is provided
	if *deploy == "" && *validate == "" {
		fmt.Println("Error: Must provide either -d or -v parameter")
		fmt.Println("Use -help for usage information")
		os.Exit(1)
	}

	// Process the provided flag
	config := config.Load()
	if *deploy != "" {
		if err := commands.DeployCommit(config, *deploy); err != nil {
			fmt.Println(err)

			os.Exit(1)
		}
	} else if *validate != "" {
		if err := commands.ValidateCommit(config, *validate); err != nil {
			fmt.Println(err)

			os.Exit(1)
		}
	}
}
