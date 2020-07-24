package jsoncase

import "encoding/json"

//TransformJSON transform the fieldName of object in json string using transformation function
func TransformJSON(transformation func(string) string) func(string) string {
	return func(j string) string {
		res := TransformJSONBytes(transformation)([]byte(j))
		return string(res)
	}
}

//TransformJSONBytes transform the fieldName of object in json bytes using transformation function
func TransformJSONBytes(transformation func(string) string) func([]byte) []byte {
	return func(j []byte) []byte {
		var m map[string]interface{}
		err := json.Unmarshal(j, &m)
		if err != nil {
			return j
		}
		res := TransformMap(transformation)(m)
		resJ, err := json.Marshal(res)
		if err != nil {
			return j
		}
		return resJ
	}
}

//ToSnakeJSON convert the fieldName of object in json string to snake_case
func ToSnakeJSON(json string) string {
	return TransformJSON(toSnakeString)(json)
}

//ToCamelJSON convert the fieldName of object in json string to camelCase
func ToCamelJSON(json string) string {
	return TransformJSON(toCamelString)(json)
}

//ToPascalJSON convert the fieldName of object in json string to PascalCase
func ToPascalJSON(json string) string {
	return TransformJSON(toPascalString)(json)
}

//ToSnakeJSONBytes convert the fieldName of object in json bytes to snake_case
func ToSnakeJSONBytes(json []byte) []byte {
	return TransformJSONBytes(toSnakeString)(json)
}

//ToCamelJSONBytes convert the fieldName of object in json bytes to camelCase
func ToCamelJSONBytes(json []byte) []byte {
	return TransformJSONBytes(toCamelString)(json)
}

//ToPascalJSONBytes convert the fieldName of object in json bytes to PascalCase
func ToPascalJSONBytes(json []byte) []byte {
	return TransformJSONBytes(toPascalString)(json)
}
