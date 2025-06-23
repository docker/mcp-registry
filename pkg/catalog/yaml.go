package catalog

import (
	"bytes"
	"os"

	"gopkg.in/yaml.v3"
)

func WriteYaml(filename string, topLevel TopLevel) error {
	var buf bytes.Buffer
	encoder := yaml.NewEncoder(&buf)
	encoder.SetIndent(2)
	if err := encoder.Encode(topLevel); err != nil {
		return err
	}

	return os.WriteFile(filename, buf.Bytes(), 0644)
}
