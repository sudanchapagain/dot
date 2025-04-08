package main

type Config struct {
	General map[string]string `toml:"general"`
	Windows map[string]string `toml:"windows"` // maybe support windows separately in future.
}

type State struct {
	Mappings map[string]string `toml:"mappings"`
}
