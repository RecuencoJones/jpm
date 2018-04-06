package main

import (
	"log"
)

// Install all dependencies defined in jpm.yml from jpm registry
func Install() {
	config := getConfig()

	createModulesDir()

	for _, dependency := range config.Dependencies {
		log.Println("Installing dependency " + dependency + "...")
		installDependency(dependency)
	}
}
