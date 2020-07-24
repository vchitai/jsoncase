package jsoncase

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStrConvert(t *testing.T) {
	testCases := []struct {
		name     string
		testFunc func(string) string
		input    string
		expected string
	}{
		{
			name:     "snake",
			testFunc: toSnakeString,
			input:    "",
			expected: "",
		},
		{
			name:     "camel",
			testFunc: toCamelString,
			input:    "",
			expected: "",
		},
		{
			name:     "pascal",
			testFunc: toPascalString,
			input:    "",
			expected: "",
		},
		{
			name:     "snake",
			testFunc: toSnakeString,
			input:    "camelCase",
			expected: "camel_case",
		},
		{
			name:     "snake",
			testFunc: toSnakeString,
			input:    "CamelCase",
			expected: "camel_case",
		},
		{
			name:     "camel",
			testFunc: toCamelString,
			input:    "camel_case",
			expected: "camelCase",
		},
		{
			name:     "camel",
			testFunc: toCamelString,
			input:    "CamelCase",
			expected: "camelCase",
		},
		{
			name:     "pascal",
			testFunc: toPascalString,
			input:    "pascal_case",
			expected: "PascalCase",
		},
		{
			name:     "pascal",
			testFunc: toPascalString,
			input:    "pascalCase",
			expected: "PascalCase",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.testFunc(tc.input))
		})
	}
}
