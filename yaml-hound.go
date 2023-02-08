// The package covering the need for non-standard approaches when working with YAML files.
package yamlhound

import (
	"errors"
)

// Struct properties UnmYAML and Footprints must be defined before the function
// is executed: UnmYAML and Footprints.
//
// The response after the function is executed will be found in the following
// properties: Caught and Found.
type YAMLTracer struct {
	UnmYAML    map[string]interface{} // unmarshalled YAML file
	Footprints []string               // the search key or strict sequence of search keys
	Caught     interface{}            // contains the value of the searched key after processing
	Found      bool                   // whether the key was found or not
}

// FootprintSniffer traverses the parsed YAML file and looks for a match against the
// supplied key. Returns the first match found.
// If more than one key is passed, the function looks for an exact match of the
// sequence in the YAML tree.
func (y *YAMLTracer) FootprintSniffer() error {
	if len(y.Footprints) < 1 {
		// There must be at least one YAML key or series of consecutive keys
		// submitted.
		return errors.New("no traces to follow.")
	}
	_, _ = y.footprintSniffer(y.UnmYAML, y.Footprints, true)

	return nil
}
