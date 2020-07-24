package jsoncase

import (
	"bytes"
	"unicode"
)

func toSnakeString(s string) string {
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

func toCamelString(s string) string {
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

func toPascalString(s string) string {
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
