package main

import (
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

func readMappings() (*Config, error) {
	mappingsPath := filepath.Join(os.Getenv("HOME"), ".dotfiles", ".mappings")
	config := &Config{}

	if _, err := toml.DecodeFile(mappingsPath, config); err != nil {
		return nil, err
	}

	return config, nil
}
