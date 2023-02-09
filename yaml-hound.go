// The package covering the need for non-standard approaches when working with YAML files.
package yamlhound

import (
	"errors"
)

// Struct properties UnmYAML and Footprints must be defined before the
// FootprintSniffer function is executed.
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

// The YAMLKeys.UnmYAML structure property must be defined before the
// FirstLevelKeys function is executed.
type YAMLKeys struct {
	UnmYAML map[string]interface{} // unmarshalled YAML file
	Caught  []string               // contains the values of the search keys by level after processing
	Found   bool                   // whether any key was found or not (it will be FALSE if the YAML document is empty)
}

// FirstLevelKeys outputs the names of the first-level keys from the parsed YAML
// file. If the file is empty YAMLKeys.Found will be FALSE, and YAMLKeys.Caught
// will be an empty slice.
func (y *YAMLKeys) FirstLevelKeys() {
	for k := range y.UnmYAML {
		y.Caught = append(y.Caught, k)
	}

	if len(y.Caught) > 0 {
		y.Found = true
	}
}
