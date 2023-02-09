package yamlhound

import (
	"fmt"
	"os"
	"testing"

	yaml "gopkg.in/yaml.v3"
)

func TestFootprintSniffer(t *testing.T) {
	if err := yamlReader("test-stuff/test.yaml", &fRead); err != nil {
		t.Errorf("Read Conf: %v", err)
		return
	}

	tests := []tfs{
		{YAMLTracer{fRead, []string{"version"}, "", true}, "Sample scenario 1 (first level)", nil},

		// Returns <nil> because no such key is found or the key is part of a
		// structure (with no specific value).
		{YAMLTracer{fRead, []string{"food"}, "", true}, nil, nil},

		{YAMLTracer{fRead, []string{"vegetables"}, "", true}, []interface{}{"tomato", "cucumber"}, nil},

		// This test case covers the scenario where we have a specified key (or
		// sequence of keys) that is contained twice in the YAML file. Since a
		// map (map[string]interface{}) is handled, the traversal sequence
		// cannot be predicted during an iteration. That is why we use the DPO
		// (Different Possible Outcomes) when checking.
		{YAMLTracer{fRead, []string{"herbivores"}, "", true}, "", []string{"deer", "cow"}},

		{YAMLTracer{fRead, []string{"config", "herbivores"}, "", true}, "deer", nil},
		{YAMLTracer{fRead, []string{"food", "herbivores"}, "", true}, "cow", nil},

		{YAMLTracer{fRead, []string{"config", "omnivores", "chickens"}, "", true}, "Brown Leghorn", nil},
		{YAMLTracer{fRead, []string{"config", "predators", "terrestrial"}, "", true}, []interface{}{"lion", "tiger", "polar bear", "wolf"}, nil},
		{YAMLTracer{fRead, []string{"predators", "amphibians"}, "", true}, "Saltwater crocodile", nil},

		{YAMLTracer{fRead, []string{"omnivores", "chickens"}, "", true}, "Brown Leghorn", nil},
		{YAMLTracer{fRead, []string{"fast-food", "chickens"}, "", true}, []interface{}{"KFC", "Chick-fil-A", "Popeyes"}, nil},
		{YAMLTracer{fRead, []string{"desserts", "chickens"}, "", true}, "Sex in a Pan", nil},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v", tc.yt.Footprints), func(t *testing.T) {
			tc.yt.FootprintSniffer()

			if tc.dpo != nil {
				for _, v := range tc.dpo {
					if v == tc.yt.Caught {
						return
					}
				}

				t.Errorf("\nwant one of these[%v] / got[%v]\n", tc.dpo, tc.yt.Caught)
			}

			if tc.expect == nil {
				if tc.yt.Caught != "" {
					t.Errorf("\nwant[%v] / got[%v]\n", "", tc.yt.Caught)
				}
			} else if fmt.Sprintf("%T", tc.yt.Caught) == "[]interface {}" {
				ci := tc.yt.Caught.([]interface{})
				ei := tc.expect.([]interface{})
				for k, v := range ci {
					if v != ei[k] {
						t.Errorf("\nwant[%v] / got[%v]\n", ei[k], v)
					}
				}
			} else {
				if tc.yt.Caught.(string) != tc.expect.(string) {
					t.Errorf("\nwant[%v] / got[%v]\n", tc.expect, tc.yt.Caught)
				}
			}
		})
	}
}

func TestFirstLevelKeys(t *testing.T) {
	yk := YAMLKeys{}
	if err := yamlReader("test-stuff/test.yaml", &yk.UnmYAML); err != nil {
		t.Errorf("Read Conf: %v", err)
		return
	}

	tc := []string{"vegetables", "version", "config", "food"}
	yk.FirstLevelKeys()

	for _, vyk := range yk.Caught {
		t.Run(fmt.Sprintf("%v", vyk), func(t *testing.T) {
			for _, vtc := range tc {
				if vyk == vtc {
					return
				}
			}
			t.Errorf("\nKey (%v) is not part of slice (%v)\n", vyk, tc)
		})
	}
}

func BenchmarkFirstLevelKeys(b *testing.B) {
	yk := YAMLKeys{}
	if err := yamlReader("test-stuff/test.yaml", &yk.UnmYAML); err != nil {
		b.Errorf("Read Conf: %v", err)
		return
	}

	b.ResetTimer()
	b.Run(fmt.Sprint("FirstLevelKeys benchmark"), func(b *testing.B) {
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			yk.FirstLevelKeys()
		}
		b.StopTimer()
	})
}

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
