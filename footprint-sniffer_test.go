package yamlhound

// import (
// 	"fmt"
// 	"os"
// 	"testing"

// 	yaml "gopkg.in/yaml.v2"
// )

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
