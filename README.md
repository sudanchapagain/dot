dot
===

A small dotfiles manager which maps your declared files to destination with
symlinks.

usage
-----

- In `~/.dotfiles/` create a new file called `.mappings`
- It should be in TOML format with everything under `[general]` table.
- The entries should be `"source" = "destination"` where source should have path
  relative to `~/.dotfiles/` whereas destination should not.

  example `.mappings`

  ```toml
  [general]
  # ghostty/config is ~/.dotfiles/ghostty/config
  "ghostty/config" = "/home/username/.config/ghostty/config"

  # starship/config is ~/.dotfiles/starship/config
  "starship/config" = "~/.config/starship/config"
  ```

- Then, start using the CLI with `dot link` or `dot l` to map everything.
- If you want to remove all links then use `dot remove` or `dot r`.
- If you want to check status of entries on which one is not linked and
  which is use `dot status` or `dot s`.

```text
USAGE:
   dotfiles [global options] command [command options]

COMMANDS:
   link, l    Link all dotfiles declared on .mappings
   remove, r  Remove all managed symlinks
   status, s  Show symlink status
   help, h    Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

**IF YOU ARE USING GIT OR ANY OTHER VERSION CONTROL REMEMBER TO IGNORE `.state`
FILE IN YOUR `~/.dotfiles/` DIRECTORY.**

building
--------

```
go build .
```

license
-------

Everything is licensed under MIT. See [LICENSE](./LICENSE)
