package yamlhound

import "testing"

type tfs struct {
	yt     YAMLTracer
	expect interface{}
	dpo    []string // Different Possible Outcomes.
}

var fRead map[string]interface{}

func TestUnmarshallCheck(t *testing.T) {
	yt := YAMLTracer{}
	if err := yt.unmarshallCheck(); err == nil {
		// The document has NOT been unmarshalled but returns a nil error.
		t.Errorf("the document has NOT been unmarshalled / got (%v)", err)
	}

	if err := yamlReader("test-stuff/test.yaml", &yt.UnmYAML); err != nil {
		t.Errorf("Read Conf: %v", err)
		return
	}
	if err := yt.unmarshallCheck(); err != nil {
		t.Errorf("the document has been unmarshalled / got (%v)", err)
	}
}

// Helper functions

func TestFootprintCheck(t *testing.T) {
	yt := YAMLTracer{}
	if err := yt.footprintsCheck(); err == nil {
		t.Errorf("not defined YAMLTracer.Footprints but no error returned / got (%v)", err)
	}

	yt.Footprints = []string{"node1", "node2"}
	if err := yt.footprintsCheck(); err != nil {
		t.Errorf("defined YAMLTracer.Footprints but error returned / got (%v)", err)
	}
}
