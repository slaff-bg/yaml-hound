package yamlhound

import (
	"fmt"
	"testing"
)

// import (
// 	"fmt"
// 	"os"
// 	"testing"

// 	yaml "gopkg.in/yaml.v2"
// )

func TestFpSniffer(t *testing.T) {
	yt := YAMLTracer{}
	// yt.Footprints = append(yt.Footprints, []string{"version"}...) // true => Sample scenario 1 (first level)
	// yt.Footprints = append(yt.Footprints, []string{"food"}...) // false => nil
	// yt.Footprints = append(yt.Footprints, []string{"vegetables"}...) // true => [tomato cucumber]

	// yt.Footprints = append(yt.Footprints, []string{"config", "herbivores"}...) // true => deer
	// yt.Footprints = append(yt.Footprints, []string{"config", "omnivores", "chickens"}...) // true => Brown Leghorns
	// yt.Footprints = append(yt.Footprints, []string{"config", "predators", "terrestrial"}...) // true => [lion tiger polar bear wolf]

	yt.Footprints = append(yt.Footprints, []string{"predators", "amphibians"}...)
	// yt.Footprints = append(yt.Footprints, []string{"omnivores", "chickens"}...)

	if err := yamlReader("test-stuff/test.yaml", &yt.UnmYAML); err != nil {
		t.Errorf("Read Conf: %v", err)
		return
	}

	yt.footprintSniffer(yt.UnmYAML, yt.Footprints, true)
	fmt.Printf("\n\nRESP ::::::: %v\n\n", yt.Caught)

	// if err := yt.FootprintSniffer(); err != nil {
	// 	t.Error(err.Error())
	// }
}

// func TestFootprintSniffer(t *testing.T) {
// 	yt := YAMLTracer{}
// 	// yt.Footprints = append(yt.Footprints, "version")
// 	// yt.Footprints = append(yt.Footprints, []string{"foo", "bar"}...)
// 	// yt.Footprints = append(yt.Footprints, []string{"fast-food", "chickens"}...)
// 	yt.Footprints = append(yt.Footprints, []string{"foo", "bar"}...)

// 	if err := yamlReader("test-stuff/test2.yaml", &yt.UnmYAML); err != nil {
// 		panic(err.Error())
// 	}
// 	if err := yt.FootprintSniffer(); err != nil {
// 		t.Error(err.Error())
// 	}

// 	// fmt.Println(yt.UnmYAML)
// 	fmt.Println(yt.Caught)
// }

// func yamlReader(f string, c *map[string]interface{}) error {
// 	yamlFile, err := os.ReadFile(f)
// 	if err != nil {
// 		return err
// 	}

// 	if err := yaml.Unmarshal([]byte(yamlFile), &c); err != nil {
// 		return err
// 	}

// 	return nil
// }
