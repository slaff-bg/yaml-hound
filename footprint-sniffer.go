package yamlhound

import "fmt"

func (y *YAMLTracer) footprintSniffer(yum map[string]interface{}, yfp []string) (map[string]interface{}, bool) {
	for yfpK, yfpV := range yfp {
		for yumK, yumV := range yum {
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

			if yumK == yfpV {
				if fmt.Sprintf("%T", yumV) == "map[string]interface {}" {
					//
				}

				if yfpK+1 == len(yfp) {
					y.Caught = yumV
					y.Found = true
					return map[string]interface{}{}, true
				}
			}
		}
	}

	return map[string]interface{}{}, false
}
