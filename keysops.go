package yamlhound

import (
	"fmt"
)

// FirstLevelKeys retrieves the first-level keys from the unmarshalled YAML
// file. If the file is empty YAMLKeys.Found will be FALSE, and YAMLKeys.Caught
// will be an empty slice.
func (y *YAMLTracer) firstLevelKeys() {
	cgth := []string{}
	for k := range y.UnmYAML {
		cgth = append(cgth, k)
	}

	if len(cgth) > 0 {
		y.Caught = cgth
		y.Found = true
	}
}

// keysOfKey loops through the unmarshalled YAML file and looks for a match
// against the provided key. Retrieves the first-level keys of the first found
// match with the provided key. If more than one key is provided, the function
// looks for an exact match of the sequence in the YAML tree.
//
//   - yum: unmarshalled YAML file
//
//   - yfp: the slice contains the search key. The part contains the search key or
//     the exact match of the sequence keys (if more than one key is provided).
//     NOTE: The key MUST NOT be an empty slice at this entry point in the codebase.
//     But it can be in the `func (y *YAMLTracer) KeysOfKey()` function.
//
//   - fid: whether to follow the index in-depth or not. (
//
//     true - trace in-depth
//     false - only trace for a matching key on the first YAML level (does not extend to sub-levels)
//
// )
func (y *YAMLTracer) keysOfKey(yum map[string]interface{}, yfp []string, fid bool) (map[string]interface{}, bool) {
	cgth := []string{}
	y.Found = false
fOut:
	for _, yfpV := range yfp {
		for yumK, yumV := range yum {
			if yfpV == yumK {
				if len(yfp)-1 == 0 {
					if fmt.Sprintf("%T", yumV) != "map[string]interface {}" {
						break fOut
					}

					for k := range yumV.(map[string]interface{}) {
						cgth = append(cgth, k)
					}
					if len(cgth) > 0 {
						y.Caught = cgth
						y.Found = true
					}
					break fOut
				}

				y.keysOfKey(yumV.(map[string]interface{}), yfp[1:], false)
				break fOut
			}

			if fid && fmt.Sprintf("%T", yumV) == "map[string]interface {}" {
				if _, ok := y.keysOfKey(yumV.(map[string]interface{}), yfp, true); ok {
					break fOut
				}
			}
		}

		break
	}

	return map[string]interface{}{}, y.Found
}
