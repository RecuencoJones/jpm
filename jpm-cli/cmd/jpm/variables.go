package main

import "github.com/RecuencoJones/jpm/util"

var (
	ConfigFileName = "jpm.yml"
	ModulesDir     = "jpm_modules"
	CurrentDir     = util.GetCwd()
	Registry       = util.GetEnv("JPM_REGISTRY", "https://registry.jpm.org")
)
