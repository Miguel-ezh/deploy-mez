package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	GitRepoPath string `yaml:"git_repo_path"`
	ProjectPath string `yaml:"project_path"`
	BackupPath  string `yaml:"backup_path"`
}

func Load() Config {
	file, err := os.Open("./config.yaml")
	if err != nil {
		log.Fatalf("Error opening config file: %v", err)
	}
	defer file.Close()

	var config Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		log.Fatalf("Error parsing config file: %v", err)
	}

	return config
}
