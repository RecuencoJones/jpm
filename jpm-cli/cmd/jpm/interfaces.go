package main

// Config describes the format of jpm configuration file
type Config struct {
	Name         string   `yaml:"string"`
	Version      string   `yaml:"version"`
	Files        string   `yaml:"files"`
	Dependencies []string `yaml:"dependencies"`
}
