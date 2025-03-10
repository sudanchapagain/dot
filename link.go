package main

import (
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

func linkDotfiles(c *cli.Context) error {
	config, err := readMappings()
	if err != nil {
		return err
	}

	if err := removeSymlinks(c); err != nil {
		return err
	}

	for source, destination := range config.General {
		sourcePath := filepath.Join(os.Getenv("HOME"), ".dotfiles", source)
		destinationPath := expandTilde(destination)

		if err := os.MkdirAll(filepath.Dir(destinationPath), 0755); err != nil {
			return err
		}

		if err := os.Symlink(sourcePath, destinationPath); err != nil {
			return err
		}

		color.Green("Linked: %s -> %s", sourcePath, destinationPath)
	}

	if err := updateState(config.General); err != nil {
		return err
	}

	return nil
}

func expandTilde(path string) string {
	if len(path) > 0 && path[0] == '~' {
		return filepath.Join(os.Getenv("HOME"), path[1:])
	}
	return path
}
