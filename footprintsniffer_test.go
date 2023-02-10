package yamlhound

import (
	"fmt"
	"testing"
)

func TestFootprintSniffer_Unit(t *testing.T) {
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
			_, _ = tc.yt.footprintSniffer(tc.yt.UnmYAML, tc.yt.Footprints, true)

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

// Benchmarks

func BenchmarkFirstLvlPrivateFunctionFootprintSniffer(b *testing.B) {
	if err := yamlReader("test-stuff/test.yaml", &fRead); err != nil {
		b.Errorf("Read Conf: %v", err)
		return
	}

	tests := []YAMLTracer{
		{fRead, []string{"version"}, "", true},
		{fRead, []string{"vegetables"}, "", true},
	}

	b.ResetTimer()
	for _, tc := range tests {
		b.Run(fmt.Sprintf("%v", tc.Footprints), func(b *testing.B) {
			b.StartTimer()
			for i := 0; i < b.N; i++ {
				_, _ = tc.footprintSniffer(tc.UnmYAML, tc.Footprints, true)
			}
			b.StopTimer()
		})
	}
}

func BenchmarkSecondLvlPrivateFunctionFootprintSniffer(b *testing.B) {
	if err := yamlReader("test-stuff/test.yaml", &fRead); err != nil {
		b.Errorf("Read Conf: %v", err)
		return
	}

	tests := []YAMLTracer{
		{fRead, []string{"herbivores"}, "", true},
		{fRead, []string{"omnivores"}, "", true},
	}

	b.ResetTimer()
	for _, tc := range tests {
		b.Run(fmt.Sprintf("%v", tc.Footprints), func(b *testing.B) {
			b.StartTimer()
			for i := 0; i < b.N; i++ {
				tc.footprintSniffer(tc.UnmYAML, tc.Footprints, true)
			}
			b.StopTimer()
		})
	}
}

func BenchmarkTwoKeysInDepthPrivateFunctionFootprintSniffer(b *testing.B) {
	if err := yamlReader("test-stuff/test.yaml", &fRead); err != nil {
		b.Errorf("Read Conf: %v", err)
		return
	}

	tests := []YAMLTracer{
		{fRead, []string{"flying", "eagles"}, "", true},
		{fRead, []string{"omnivores", "chickens"}, "", true},
	}

	b.ResetTimer()
	for _, tc := range tests {
		b.Run(fmt.Sprintf("%v", tc.Footprints), func(b *testing.B) {
			b.StartTimer()
			for i := 0; i < b.N; i++ {
				tc.footprintSniffer(tc.UnmYAML, tc.Footprints, true)
			}
			b.StopTimer()
		})
	}
}
