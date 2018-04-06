package main

func Add(dependency string) {
	config := getConfig()
	newDeps := make([]string, len(config.Dependencies)+1)

	createModulesDir()

	installDependency(dependency)

	newDeps = append(config.Dependencies, dependency)
	config.Dependencies = newDeps

	writeConfig(config)
}
