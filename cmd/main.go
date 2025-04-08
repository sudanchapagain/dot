package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:    "dotfiles",
		Usage:   "Manage dotfiles with symlinks",
		Version: "1.0.0",
		Commands: []*cli.Command{
			{
				Name:    "link",
				Aliases: []string{"l"},
				Usage:   "Link all dotfiles declared on .mappings",
				Action:  linkDotfiles,
			},
			{
				Name:    "remove",
				Aliases: []string{"r"},
				Usage:   "Remove all managed symlinks",
				Action:  removeSymlinks,
			},
			{
				Name:    "status",
				Aliases: []string{"s"},
				Usage:   "Show symlink status",
				Action:  showStatus,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
