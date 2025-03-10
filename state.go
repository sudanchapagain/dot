package main

import (
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

func readState() (*State, error) {
	statePath := filepath.Join(os.Getenv("HOME"), ".dotfiles", ".state")
	state := &State{}

	if _, err := toml.DecodeFile(statePath, state); err != nil {
		if !os.IsNotExist(err) {
			return nil, err
		}
	}

	return state, nil
}

func updateState(mappings map[string]string) error {
	statePath := filepath.Join(os.Getenv("HOME"), ".dotfiles", ".state")
	state := &State{Mappings: mappings}

	file, err := os.Create(statePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := toml.NewEncoder(file)
	return encoder.Encode(state)
}
