package yamlhound

import (
	"fmt"
)

// ti - trace index (
//
//	true - trace indepth;
//	false - trace for succes only first lvl keys (yum);
//
// )
func (y *YAMLTracer) footprintSniffer(yum map[string]interface{}, yfp []string, ti bool) (map[string]interface{}, bool) {
fOut:
	for _, yfpV := range yfp {
		for yumK, yumV := range yum {
			fmt.Printf("\n::: yumK (%v) :::\n", yumK)
			// 1. The first key searched matches the YAML crawl key.
			// 		1.1. We check for the type of value.
			//          1.1.1. The value is different from the MAP (we return it because it is assumed that it's the value searched for).
			//          1.1.2. The value is MAP (do recursion).
			//              -- remove one element from the searched keys, and if there is another key, we continue cycling.

			//  1.1.1. Finds a first-level key.
			if yfpV == yumK {
				if fmt.Sprintf("%T", yumV) != "map[string]interface {}" {
					y.Caught, y.Found = yumV, true
					break fOut
				}

				if len(yfp)-1 > 0 {
					if _, ok := y.footprintSniffer(yumV.(map[string]interface{}), yfp[1:], false); ok {
						y.Found = true
					}
				}

				break fOut
			}

			//  1.1.2.
			// if _, ok := y.footprintSniffer(yumV.(map[string]interface{}), yfp, true); ok {
			// 	y.Found = true
			// 	break fOut
			// }

			// if fmt.Sprintf("%T", yumV) == "map[string]interface {}" {
			// if yumK == yfpV {

			// }
			// res, ok := y.footprintSniffer(yumV.(map[string]interface{}), yfp[0:])
			// if ok {
			// 	break
			// 	return res, true
			// }
			// return map[string]interface{}{}, false
			// continue
			// }

			// if yumK == yfpV {
			// 	if yfpK+1 == len(yfp) {
			// 		y.Caught = yumV
			// 		y.Found = true
			// 		return map[string]interface{}{}, true
			// 	}
			// }

			// return map[string]interface{}{}, false

			// if yumK == yfpV {
			// 	if fmt.Sprintf("%T", yumV) == "map[string]interface {}" {
			// 		//
			// 	}

			// 	if yfpK+1 == len(yfp) {
			// 		y.Caught = yumV
			// 		y.Found = true
			// 		// return map[string]interface{}{}, true
			// 	}
			// }
		}
	}

	return map[string]interface{}{}, y.Found
}
