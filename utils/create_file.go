package utils

import (
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v2"
)

// CreateNewFile allows to create new file with a given name
func CreateNewFile(fileName string) (*os.File, error) {
	// create new file with given name
	file, err := os.Create(fileName)
	if err != nil {
		return nil, fmt.Errorf("error creating file: %v", err)
	}

	return file, nil
}

// SerializeToYAML marshals the config data into YAML format
// and writes it to the provided writer
func SerializeToYAML(writer io.Writer, cfg interface{}) error {
	encoder := yaml.NewEncoder(writer)
	defer encoder.Close()
	return encoder.Encode(cfg)
}
