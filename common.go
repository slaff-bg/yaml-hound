package yamlhound

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// unmarshallCheck checks if the YAML file is not unmarshalled.
func (y *YAMLTracer) unmarshallCheck() error {
	if fmt.Sprintf("%T", y.UnmYAML) != "map[string]interface {}" {
		// - or you are probably using an older version of the gopkg.in/yaml.v3
		// package
		return errors.New("unmarshalization failure")
	}
	// - an empty YAML file
	if len(y.UnmYAML) == 0 {
		return errors.New("unmarshalization failure or an empty YAML file")
	}
	return nil
}

// footprintsCheck cheks if the YAMLTracer.Footprints are not defined. There
// must be at least one YAML key or series of consecutive keys provided.
func (y *YAMLTracer) footprintsCheck() error {
	if len(y.Footprints) < 1 {
		return errors.New("no traces to follow")
	}
	return nil
}

// Helper functions

// yamlReader is a helper function for the execution of the tests.
func yamlReader(f string, c *map[string]interface{}) error {
	yamlFile, err := os.ReadFile(f)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal([]byte(yamlFile), &c); err != nil {
		return err
	}

	return nil
}
