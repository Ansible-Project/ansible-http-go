package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Port              int    `yaml:"port"`
	WorkDir           string `yaml:"work_dir"`
	Mode              string `yaml:"mode"`
	RepositoryUrl     string `yaml:"repository_url"` 
	DefaultInventory  string `yaml:"default_inventory"`
	DefaultVerbose    string `yaml:"default_verbose"`
	DefaultBranch     string `yaml:"default_branch"`
}

func loadConfig(configPath string) (*Config, error) {
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("Failed to load config: %s, err:[%s]", configPath, err)
	}

	var config *Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("Failed to parse yaml. err:[%s]", err)
	}
	
	config = initializeConfig(config)
	
	err = validateConfig(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func initializeConfig(config *Config) *Config {
	if config.WorkDir == "" {
		config.WorkDir = "work"
	}
	if config.Mode == "" {
		config.Mode = "git"
	}
	return config
}

func validateConfig(config *Config) error {
	if config.Port < 1 || config.Port > 65535 {
		return fmt.Errorf("Invalid port: %d", config.Port)
	}
	if config.Mode != "" && config.Mode != "git" {
		return fmt.Errorf("Invalid mode: %s", config.Mode)
	}
	if config.DefaultVerbose != "" && !validateVerbose(config.DefaultVerbose) {
		return fmt.Errorf("Invalid default_verbose: %s", config.DefaultVerbose)
	}
	dir, err := os.Stat(config.WorkDir)
	if err != nil {
		return fmt.Errorf("Not exists: %s", config.WorkDir)
	}
	if !dir.IsDir() {
		return fmt.Errorf("Not directory: %s", config.WorkDir)
	}

	return nil
}

func validateVerbose(verbose string) bool {
	if verbose == "-v" {
		return true
	}
	if verbose == "-vv" {
		return true
	}
	if verbose == "-vvv" {
		return true
	}
	if verbose == "-vvvv" {
		return true
	}
	return false
}
