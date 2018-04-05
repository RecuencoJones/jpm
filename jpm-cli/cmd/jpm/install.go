package main

import (
	"log"
	"os"
	"path"
	"strings"

	"github.com/RecuencoJones/jpm/util"
)

var (
	configFileName = "jpm.yml"
	modulesDir     = "jpm_modules"
	currentDir     = util.GetCwd()
	registry       = util.GetEnv("JPM_REGISTRY", "https://registry.jpm.org")
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

func getConfig() Config {
	config := Config{}

	err := util.ReadYAML(path.Join(currentDir, configFileName), &config)
	if err != nil {
		log.Fatal(err)
	}

	return config
}

func createModulesDir() {
	_ = os.Mkdir(path.Join(currentDir, modulesDir), 0777)
}

func installDependency(dependency string) {
	name, version := parseDependency(dependency)

	err := fetchDependency(name, version)

	if err != nil {
		log.Fatal(err)
	}
}

func parseDependency(dependency string) (string, string) {
	parts := strings.Split(dependency, "@")

	name := parts[0]
	version := "latest"

	if len(parts) == 2 {
		version = parts[1]
	}

	return name, version
}

func fetchDependency(name, version string) error {
	os.Mkdir(path.Join(currentDir, modulesDir, name), 0777)

	log.Println("http " + registry + "/package/" + name + "@" + version)

	return nil
}
