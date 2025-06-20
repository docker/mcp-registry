package servers

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func Read(path string) (Server, error) {
	file, err := os.Open(path)
	if err != nil {
		return Server{}, err
	}
	defer file.Close()

	var server Server
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&server); err != nil {
		return Server{}, fmt.Errorf("failed to decode server file %s: %w", path, err)
	}

	return server, nil
}
