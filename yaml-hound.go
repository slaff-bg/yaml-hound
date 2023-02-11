// The package covers uncommon needs when working with YAML files.
package yamlhound

// Struct properties UnmYAML and Footprints must be defined before the
// FootprintSniffer function is executed.
//
// The response after the function is executed will be found in the following
// properties: Caught and Found.
type YAMLTracer struct {
	UnmYAML    map[string]interface{} // unmarshalled YAML file
	Footprints []string               // the search key or strict sequence of search keys
	Caught     interface{}            // contains the execution result
	Found      bool                   // whether there is any result of executing the function or not
}

// FootprintSniffer retrives the first found match against the provided key.
// If more than one key is provided, the function looks for an exact match with
// the provided set of keys.
func (y *YAMLTracer) FootprintSniffer() error {
	if err := y.unmarshallCheck(); err != nil {
		return err
	}
	if err := y.footprintsCheck(); err != nil {
		return err
	}
	_, _ = y.footprintSniffer(y.UnmYAML, y.Footprints, true)

	return nil
}

// KeysOfKey retrieves the first-level keys of the first found match against the
// provided key. If more than one key is provided, the function looks for an
// exact match of the sequence in the YAML tree.
//
// If no YAMLTracer.Footprints are provided will collect the first-level keys of
// the YAML structure.
func (y *YAMLTracer) KeysOfKey() error {
	if err := y.unmarshallCheck(); err != nil {
		return err
	}

	if err := y.footprintsCheck(); err != nil {
		y.firstLevelKeys()
	} else {
		_, _ = y.keysOfKey(y.UnmYAML, y.Footprints, true)
	}

	return nil
}
