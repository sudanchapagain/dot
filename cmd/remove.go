package main

import (
	"os"

	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

func removeSymlinks(c *cli.Context) error {
	state, err := readState()
	if err != nil {
		return err
	}

	for _, destination := range state.Mappings {
		destinationPath := expandTilde(destination)

		if err := os.Remove(destinationPath); err != nil {
			if !os.IsNotExist(err) {
				return err
			}
		}

		color.Red("Removed: %s", destinationPath)
	}

	if err := updateState(make(map[string]string)); err != nil {
		return err
	}

	return nil
}
