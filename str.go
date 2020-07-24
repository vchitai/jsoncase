package jsoncase

import "encoding/json"

func TransformJson(transformation func(string) string) func(string) string {
	return func(j string) string {
		res := TransformJsonBytes(transformation)([]byte(j))
		return string(res)
	}
}

func TransformJsonBytes(transformation func(string) string) func([]byte) []byte {
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

func ToSnakeJson(json string) string {
	return TransformJson(ToSnakeString)(json)
}

func ToCamelJson(json string) string {
	return TransformJson(ToCamelString)(json)
}

func ToPascalJson(json string) string {
	return TransformJson(ToPascalString)(json)
}

func ToSnakeJsonBytes(json []byte) []byte {
	return TransformJsonBytes(ToSnakeString)(json)
}

func ToCamelJsonBytes(json []byte) []byte {
	return TransformJsonBytes(ToCamelString)(json)
}

func ToPascalJsonBytes(json []byte) []byte {
	return TransformJsonBytes(ToPascalString)(json)
}
