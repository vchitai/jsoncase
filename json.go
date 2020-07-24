package jsoncase

import (
	"bytes"
	"unicode"
)

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
				thisVal[idx] = TransformFieldName(elem.(map[string]interface{}))
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

func ToSnakeString(s string) string {
	if len(s) == 0 {
		return s
	}
	var buffer bytes.Buffer
	for _, c := range s {
		if unicode.IsUpper(c) {
			if buffer.Len() > 0 {
				buffer.WriteRune('_')
			}
			buffer.WriteRune(unicode.ToLower(c))
		} else {
			buffer.WriteRune(c)
		}
	}
	return buffer.String()
}

func ToCamelString(s string) string {
	if len(s) == 0 {
		return s
	}
	var buffer bytes.Buffer
	var skip bool
	for _, c := range s {
		if c == '_' {
			skip = true
			continue
		}
		if skip {
			buffer.WriteRune(unicode.ToUpper(c))
		} else if buffer.Len() == 0 {
			buffer.WriteRune(unicode.ToLower(c))
		} else {
			buffer.WriteRune(c)
		}
	}
	return buffer.String()
}

func ToPascalString(s string) string {
	if len(s) == 0 {
		return s
	}
	var buffer bytes.Buffer
	var skip bool
	for _, c := range s {
		if c == '_' {
			skip = true
			continue
		}
		if skip || buffer.Len() == 0 {
			buffer.WriteRune(unicode.ToUpper(c))
		} else {
			buffer.WriteRune(c)
		}
	}
	return buffer.String()
}

func ToSnakeVal(json interface{}) interface{} {
	return TransformVal(ToSnakeString)(json)
}

func ToCamelVal(json interface{}) interface{} {
	return TransformVal(ToCamelString)(json)
}

func ToPascalVal(json interface{}) interface{} {
	return TransformVal(ToPascalString)(json)
}
