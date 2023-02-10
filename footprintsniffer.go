package yamlhound

import (
	"fmt"
)

// The function traverses the unmarshalled YAML file and looks for a match
// against the supplied key. Returns the first match found.
// If more than one key is passed, the function looks for an exact match of the
// sequence in the YAML tree.
//
// yum - unmarshalled YAML file
// yfp - the slice contains the search key. The part contains the search key or
// the exact match of the sequence keys (if more than one key is supplied).
// fid - whether to follow the index in-depth or not. (
//
//	true - trace in-depth
//	false - only trace for a matching key on the first YAML level (does not extend to sub-levels)
//
// )
func (y *YAMLTracer) footprintSniffer(yum map[string]interface{}, yfp []string, fid bool) (map[string]interface{}, bool) {
	y.Found = false
fOut:
	for _, yfpV := range yfp {
		for yumK, yumV := range yum {
			if yfpV == yumK {
				if fmt.Sprintf("%T", yumV) != "map[string]interface {}" {
					y.Caught, y.Found = yumV, true
					break fOut
				}

				if len(yfp)-1 > 0 {
					y.footprintSniffer(yumV.(map[string]interface{}), yfp[1:], false)
				}

				break fOut
			}

			if fid && fmt.Sprintf("%T", yumV) == "map[string]interface {}" {
				if _, ok := y.footprintSniffer(yumV.(map[string]interface{}), yfp, true); ok {
					break fOut
				}
			}
		}

		break fOut
	}

	return map[string]interface{}{}, y.Found
}
