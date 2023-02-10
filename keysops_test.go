package yamlhound

import (
	"fmt"
	"testing"
)

func TestFirstLevelKeys_Unit(t *testing.T) {
	yt := YAMLTracer{}
	if err := yamlReader("test-stuff/test.yaml", &yt.UnmYAML); err != nil {
		t.Errorf("Read Conf: %v", err)
		return
	}

	tc := []string{"vegetables", "version", "config", "food"}
	yt.firstLevelKeys()
	for _, vyt := range yt.Caught.([]string) {
		t.Run(fmt.Sprintf("TestFirstLevelKeys_Unit: %v", vyt), func(t *testing.T) {
			for _, vtc := range tc {
				if vyt == vtc {
					return
				}
			}
			t.Errorf("\nKey (%v) is not part of slice (%v)\n", vyt, tc)
		})
	}
}

func TestKeysOfKey_Unit(t *testing.T) {
	if err := yamlReader("test-stuff/test.yaml", &fRead); err != nil {
		t.Errorf("Read Conf: %v", err)
		return
	}

	tests := []tfs{
		// First-level key NO sub-level.
		{YAMLTracer{fRead, []string{"version"}, nil, true}, nil, nil},

		// First-level key WITH sub-level.
		{YAMLTracer{fRead, []string{"food"}, nil, true}, []string{"fast-food", "desserts", "herbivores", "omnivores"}, nil},

		// Second-level key WITH sub-level.
		{YAMLTracer{fRead, []string{"predators"}, nil, true}, []string{"amphibians", "terrestrial", "flying"}, nil},

		// Second-level key WITH sub-level by using a strict key sequence.
		{YAMLTracer{fRead, []string{"config", "predators"}, nil, true}, []string{"amphibians", "terrestrial", "flying"}, nil},

		// Third-level key WITH sub-level by using a strict key sequence.
		{YAMLTracer{fRead, []string{"predators", "flying"}, nil, true}, []string{"eagles", "vultures", "owls"}, nil},

		// This test case covers the scenario where we have a specified key (or
		// sequence of keys) that is contained twice in the YAML file. Since a
		// map (map[string]interface{}) is handled, the traversal sequence
		// cannot be predicted during an iteration. That is why we use the DPO
		// (Different Possible Outcomes) when checking.
		//
		// Duplicate keys from different levels With sub-levels.
		// {YAMLTracer{fRead, []string{"pigs"}, "", true}, []string{"European", "Asians", "American"}, []string{"fried", "grilled", "in-oven"}},
		{YAMLTracer{fRead, []string{"pigs"}, "", true}, []string{"grilled", "in-oven"}, []string{"European", "Asians", "American"}},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v", tc.yt.Footprints), func(t *testing.T) {
			_, _ = tc.yt.keysOfKey(tc.yt.UnmYAML, tc.yt.Footprints, true)

			// Different Possible Outcomes (DPO) scenario.
			if tc.dpo != nil {
				if len(tc.yt.Caught.([]string)) != len(tc.expect.([]string)) && len(tc.yt.Caught.([]string)) != len(tc.dpo) {
					t.Errorf("\nwant exp.len[%v] or dpo.len[%v] / got len [%v]\n", len(tc.expect.([]string)), len(tc.dpo), len(tc.yt.Caught.([]string)))
				}
			fOut:
				for _, vyt := range tc.yt.Caught.([]string) {
					for _, vtc := range tc.expect.([]string) {
						if vyt == vtc {
							continue fOut
						}
					}
					for _, vtc := range tc.dpo {
						if vyt == vtc {
							continue fOut
						}
					}
					t.Errorf("\nKey (%v) is not part of slice (%v)\n", vyt, tc.expect.([]string))
				}
				return
			}

			// Common scenario.
			if tc.expect == nil {
				if fmt.Sprintf("%T", tc.yt.Caught) == "[]string" {
					t.Errorf("\nwant [%T] / got (%T)\n", tc.expect, tc.yt.Caught)
				}
			} else {
				if len(tc.yt.Caught.([]string)) != len(tc.expect.([]string)) {
					t.Errorf("\nwant len [%v] / got len [%v]\n", len(tc.expect.([]string)), len(tc.yt.Caught.([]string)))
				}
			fOut_:
				for _, vyt := range tc.yt.Caught.([]string) {
					for _, vtc := range tc.expect.([]string) {
						if vyt == vtc {
							continue fOut_
						}
					}
					t.Errorf("\nKey (%v) is not part of slice (%v)\n", vyt, tc.expect.([]string))
				}
			}
		})
	}
}

// - - - - - - -
// Benchmarks
// - - - - - - -

func BenchmarkFirstLevelKeys_Unit(b *testing.B) {
	yt := YAMLTracer{}
	if err := yamlReader("test-stuff/test.yaml", &yt.UnmYAML); err != nil {
		b.Errorf("Read Conf: %v", err)
		return
	}

	b.ResetTimer()
	b.Run(fmt.Sprint("FirstLevelKeys benchmark"), func(b *testing.B) {
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			yt.firstLevelKeys()
		}
		b.StopTimer()
	})
}
