package jsoncase

//TransformVal transform all fieldName of map[string]interface embedded in the interface{} using the transformation function
func TransformVal(transformation func(string) string) func(json interface{}) interface{} {
	var TransformFieldName func(val interface{}) interface{}
	TransformFieldName = func(val interface{}) interface{} {
		switch val.(type) {
		case map[string]interface{}:
			realVal := val.(map[string]interface{})
			thisVal := make(map[string]interface{}, len(realVal))
			for key, value := range realVal {
				thisVal[transformation(key)] = TransformFieldName(value)
			}
			return thisVal
		case []interface{}:
			realVal := val.([]interface{})
			thisVal := make([]interface{}, len(realVal))
			for idx, elem := range realVal {
				thisVal[idx] = TransformFieldName(elem)
			}
			return thisVal
		case []map[string]interface{}:
			realVal := val.([]map[string]interface{})
			thisVal := make([]map[string]interface{}, len(realVal))
			for idx, elem := range realVal {
				thisVal[idx] = TransformFieldName(elem).(map[string]interface{})
			}
			return thisVal
		default:
			return val
		}
	}
	return TransformFieldName
}

//TransformMap transform fieldName of all map[string]interface embedded in the root map using the transformation function
func TransformMap(transformation func(string) string) func(json map[string]interface{}) map[string]interface{} {
	var TransformFieldName func(json map[string]interface{}) map[string]interface{}
	TransformFieldName = func(json map[string]interface{}) map[string]interface{} {
		res := make(map[string]interface{}, len(json))
		for key, val := range json {
			newVal := val
			switch val.(type) {
			case map[string]interface{}:
				realVal := val.(map[string]interface{})
				newVal = TransformFieldName(realVal)
			case []interface{}:
				realVal := val.([]interface{})
				thisVal := make([]interface{}, len(realVal))
				for idx, elem := range realVal {
					switch elem.(type) {
					case map[string]interface{}:
						thisVal[idx] = TransformFieldName(elem.(map[string]interface{}))
					default:
						thisVal[idx] = elem
					}
				}
				newVal = thisVal
			case []map[string]interface{}:
				realVal := val.([]map[string]interface{})
				thisVal := make([]map[string]interface{}, len(realVal))
				for idx, elem := range realVal {
					thisVal[idx] = TransformFieldName(elem)
				}
				newVal = thisVal
			}
			res[transformation(key)] = newVal
		}
		return res
	}
	return TransformFieldName
}

//ToSnakeVal transform all fieldName of map[string]interface embedded in the interface{} to snake_case
func ToSnakeVal(json interface{}) interface{} {
	return TransformVal(toSnakeString)(json)
}

//ToCamelVal transform all fieldName of map[string]interface embedded in the interface{} to camelCase
func ToCamelVal(json interface{}) interface{} {
	return TransformVal(toCamelString)(json)
}

//ToPascalVal transform all fieldName of map[string]interface embedded in the interface{} to PascalCase
func ToPascalVal(json interface{}) interface{} {
	return TransformVal(toPascalString)(json)
}
