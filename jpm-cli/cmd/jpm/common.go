package main

import (
	"log"
	"os"
	"path"
	"strings"

	"github.com/RecuencoJones/jpm/util"
)

func getConfig() Config {
	config := Config{}

	err := util.ReadYAML(path.Join(CurrentDir, ConfigFileName), &config)
	if err != nil {
		log.Fatal(err)
	}

	return config
}

func writeConfig(c Config) {
	err := util.WriteYAML(path.Join(CurrentDir, ConfigFileName), c)
	if err != nil {
		log.Fatal(err)
	}
}

func createModulesDir() {
	_ = os.Mkdir(path.Join(CurrentDir, ModulesDir), 0777)
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
	os.Mkdir(path.Join(CurrentDir, ModulesDir, name), 0777)

	log.Println("http " + Registry + "/package/" + name + "@" + version)

	return nil
}
