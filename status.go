package main

import (
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

func showStatus(c *cli.Context) error {
	config, err := readMappings()
	if err != nil {
		return err
	}

	state, err := readState()
	if err != nil {
		return err
	}

	for source, destination := range config.General {
		sourcePath := filepath.Join(os.Getenv("HOME"), ".dotfiles", source)
		destinationPath := expandTilde(destination)

		if link, err := os.Readlink(destinationPath); err == nil {
			if link == sourcePath {
				color.Green("OK: %s -> %s", sourcePath, destinationPath)
			} else {
				color.Red("MISMATCH: %s -> %s (expected %s)", destinationPath, link, sourcePath)
			}
		} else if os.IsNotExist(err) {
			color.Yellow("MISSING: %s -> %s", sourcePath, destinationPath)
		} else {
			return err
		}
	}

	for source, destination := range state.Mappings {
		if _, exists := config.General[source]; !exists {
			destinationPath := expandTilde(destination)

			if link, err := os.Readlink(destinationPath); err == nil {
				color.Red("ORPHANED: %s -> %s (will be removed on 'link')", destinationPath, link)
			} else if os.IsNotExist(err) {
				color.Yellow("ORPHANED (MISSING): %s (not in .mappings)", destinationPath)
			} else {
				return err
			}
		}
	}

	return nil
}
