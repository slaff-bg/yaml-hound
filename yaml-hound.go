package yamlhound

import (
	"errors"
	"fmt"
)

type YAMLTracer struct {
	UnmYAML    map[string]interface{}
	Footprints []string
	Caught     interface{}
	Found      bool
}

func (y *YAMLTracer) FootprintSniffer() error {
	fl := len(y.Footprints)
	if fl < 1 {
		// There must be at least one YAML key or series of consecutive keys
		// submitted.
		return errors.New("no traces to follow.")
	}

	_, _ = y.footprintSniffer(y.UnmYAML, y.Footprints)

	fmt.Println("\n::::::::::", y.Caught, "::::::::::")

	return nil
}
