package util

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// ReadYAML reads a file and maps its contents to a given model struct
func ReadYAML(path string, model interface{}) error {
	yamlFile, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Print(err)
	}

	return yaml.Unmarshal(yamlFile, model)
}
