package jsoncase

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTransformString(t *testing.T) {
	snakeJSON := `{"a_b":{"c_d":"e"},"b_e":"x","e_f":[{"c_f":"e"}],"x_y":{"a_b":{"c_d":"e"},"b_e":"x","e_f":[{"c_f":"e"}]}}`
	camelJSON := `{"aB":{"cD":"e"},"bE":"x","eF":[{"cF":"e"}],"xY":{"aB":{"cD":"e"},"bE":"x","eF":[{"cF":"e"}]}}`
	pascalJSON := `{"AB":{"CD":"e"},"BE":"x","EF":[{"CF":"e"}],"XY":{"AB":{"CD":"e"},"BE":"x","EF":[{"CF":"e"}]}}`

	testCases := []struct {
		name          string
		transformFunc func(json string) string
		input         string
		expected      string
	}{
		{
			name:          "camel-snake",
			transformFunc: ToSnakeJSON,
			input:         camelJSON,
			expected:      snakeJSON,
		},
		{
			name:          "pascal-snake",
			transformFunc: ToSnakeJSON,
			input:         pascalJSON,
			expected:      snakeJSON,
		},
		{
			name:          "snake-camel",
			transformFunc: ToCamelJSON,
			input:         snakeJSON,
			expected:      camelJSON,
		},
		{
			name:          "pascal-camel",
			transformFunc: ToCamelJSON,
			input:         pascalJSON,
			expected:      camelJSON,
		},
		{
			name:          "camel-pascal",
			transformFunc: ToPascalJSON,
			input:         camelJSON,
			expected:      pascalJSON,
		},
		{
			name:          "snake-pascal",
			transformFunc: ToPascalJSON,
			input:         snakeJSON,
			expected:      pascalJSON,
		},
		{
			name: "error json",
			transformFunc: TransformJSON(func(s string) string {
				return s
			}),
			input:    "this is not a json string",
			expected: "this is not a json string",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.transformFunc(tc.input)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestTransformBytes(t *testing.T) {
	snakeJSON := []byte(`{"a_b":{"c_d":"e"},"b_e":"x","e_f":[{"c_f":"e"}],"x_y":{"a_b":{"c_d":"e"},"b_e":"x","e_f":[{"c_f":"e"}]}}`)
	camelJSON := []byte(`{"aB":{"cD":"e"},"bE":"x","eF":[{"cF":"e"}],"xY":{"aB":{"cD":"e"},"bE":"x","eF":[{"cF":"e"}]}}`)
	pascalJSON := []byte(`{"AB":{"CD":"e"},"BE":"x","EF":[{"CF":"e"}],"XY":{"AB":{"CD":"e"},"BE":"x","EF":[{"CF":"e"}]}}`)

	testCases := []struct {
		name          string
		transformFunc func(json []byte) []byte
		input         []byte
		expected      []byte
	}{
		{
			name:          "camel-snake",
			transformFunc: ToSnakeJSONBytes,
			input:         camelJSON,
			expected:      snakeJSON,
		},
		{
			name:          "pascal-snake",
			transformFunc: ToSnakeJSONBytes,
			input:         pascalJSON,
			expected:      snakeJSON,
		},
		{
			name:          "snake-camel",
			transformFunc: ToCamelJSONBytes,
			input:         snakeJSON,
			expected:      camelJSON,
		},
		{
			name:          "pascal-camel",
			transformFunc: ToCamelJSONBytes,
			input:         pascalJSON,
			expected:      camelJSON,
		},
		{
			name:          "camel-pascal",
			transformFunc: ToPascalJSONBytes,
			input:         camelJSON,
			expected:      pascalJSON,
		},
		{
			name:          "snake-pascal",
			transformFunc: ToPascalJSONBytes,
			input:         snakeJSON,
			expected:      pascalJSON,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.transformFunc(tc.input)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
