package mapFlattener

import "fmt"

func If(condition bool, trueVal, falseVal string) string {
	if condition {
		return trueVal
	}
	return falseVal
}

func Flatten(inMap map[string]interface{}) (outMap map[string]interface{}) {

	outMap = make(map[string]interface{})

	var parseMap func(path string, aMap map[string]interface{})
	//var parseArray func(path string, anArray []interface{})
	var parseArray_keep_list func(path string, anArray []interface{})

	parseMap =
		func(path string, aMap map[string]interface{}) {
			for key, val := range aMap {
				switch concreteVal := val.(type) {
				case map[string]interface{}:
					//fmt.Println(path)
					tmp_path := If(len(path) == 0, key, path+"."+key)
					parseMap(tmp_path, val.(map[string]interface{}))
				case []interface{}:
					//fmt.Println(path)
					tmp_path := If(len(path) == 0, key, path+"."+key)
					parseArray_keep_list(tmp_path, val.([]interface{}))
				default:
					tmp_path := If(len(path) == 0, key, path+"."+key)
					outMap[tmp_path] = concreteVal
					fmt.Println(tmp_path, ":", concreteVal)
				}
				//fmt.Println(path, ":", val)
			}
		}
	parseArray_keep_list =
		func(path string, anArray []interface{}) {
			outMap[path] = anArray
		}

	parseMap("", inMap)
	return outMap
}
