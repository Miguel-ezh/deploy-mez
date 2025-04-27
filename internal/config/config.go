package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Backup struct {
	Enabled    bool   `yaml:"enabled"`
	BackupPath string `yaml:"backup_path"`
}

type Project struct {
	NormalizedPath string `yaml:"normalized_path"`
	ProjectPath    string `yaml:"project_path"`
}

type Config struct {
	GitRepoPath string  `yaml:"git_repo_path"`
	Project     Project `yaml:"project"`
	Backup      Backup  `yaml:"backup"`
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
