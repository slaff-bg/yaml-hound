package yamlhound

import (
	"os"
	"testing"

	yaml "gopkg.in/yaml.v3"
)

func TestFootprintSniffer(t *testing.T) {
	yt := YAMLTracer{}
	// yt.Footprints = append(yt.Footprints, "version")
	// yt.Footprints = append(yt.Footprints, []string{"omnivores", "chickens"}...)
	// yt.Footprints = append(yt.Footprints, []string{"fast-food", "chickens"}...)
	// yt.Footprints = append(yt.Footprints, []string{"foo", "bar"}...)
	// test2.yaml
	// yt.Footprints = append(yt.Footprints, []string{"third", "fourth"}...)
	// yt.Footprints = append(yt.Footprints, []string{"second", "third", "fourth"}...)

	if err := yamlReader("test-stuff/test.yaml", &yt.UnmYAML); err != nil {
		t.Errorf("Read Conf: %v", err)
		return
	}

	// yt.footprintSniffer(yt.UnmYAML, []string{"second", "third", "fourth"})
	// fmt.Printf("\n\nRESP ::::::: %v\n\n", yt.Caught)

	if err := yt.FootprintSniffer(); err != nil {
		t.Error(err.Error())
	}
	// fmt.Printf("\n\nRESP ::::::: %v\n\n", yt.Caught)

	// fmt.Println(yt.UnmYAML)
	// fmt.Println(yt.Caught)
}

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
