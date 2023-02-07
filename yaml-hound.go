package yamlhound

import (
	"errors"
)

// Struct properties UnmYAML and Footprints must be defined before the function
// is executed.
//
// - UnmYAML: unmarshalled YAML file
// - Footprints: the search key or strict sequence of search keys
//
// The response after the function is executed will be found in the following
// properties:
// - Caught: contains the value of the searched key after processing
// - Found: whether the key was found or not
type YAMLTracer struct {
	UnmYAML    map[string]interface{}
	Footprints []string
	Caught     interface{}
	Found      bool
}

// The function traverses the parsed YAML file and looks for a match against the
// supplied key. Returns the first match found.
// If more than one key is passed, the function looks for an exact match of the
// sequence in the YAML tree.
func (y *YAMLTracer) FootprintSniffer() error {
	fl := len(y.Footprints)
	if fl < 1 {
		// There must be at least one YAML key or series of consecutive keys
		// submitted.
		return errors.New("no traces to follow.")
	}
	_, _ = y.footprintSniffer(y.UnmYAML, y.Footprints, true)

	return nil
}
