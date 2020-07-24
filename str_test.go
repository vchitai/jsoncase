package jsoncase

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTransformString(t *testing.T) {
	snakeJson := `{"a_b":{"c_d":"e"},"b_e":"x","e_f":[{"c_f":"e"}],"x_y":{"a_b":{"c_d":"e"},"b_e":"x","e_f":[{"c_f":"e"}]}}`
	camelJson := `{"aB":{"cD":"e"},"bE":"x","eF":[{"cF":"e"}],"xY":{"aB":{"cD":"e"},"bE":"x","eF":[{"cF":"e"}]}}`
	pascalJson := `{"AB":{"CD":"e"},"BE":"x","EF":[{"CF":"e"}],"XY":{"AB":{"CD":"e"},"BE":"x","EF":[{"CF":"e"}]}}`

	testCases := []struct {
		name          string
		transformFunc func(json string) string
		input         string
		expected      string
	}{
		{
			name:          "camel-snake",
			transformFunc: ToSnakeJson,
			input:         camelJson,
			expected:      snakeJson,
		},
		{
			name:          "pascal-snake",
			transformFunc: ToSnakeJson,
			input:         pascalJson,
			expected:      snakeJson,
		},
		{
			name:          "snake-camel",
			transformFunc: ToCamelJson,
			input:         snakeJson,
			expected:      camelJson,
		},
		{
			name:          "pascal-camel",
			transformFunc: ToCamelJson,
			input:         pascalJson,
			expected:      camelJson,
		},
		{
			name:          "camel-pascal",
			transformFunc: ToPascalJson,
			input:         camelJson,
			expected:      pascalJson,
		},
		{
			name:          "snake-pascal",
			transformFunc: ToPascalJson,
			input:         snakeJson,
			expected:      pascalJson,
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
	snakeJson := []byte(`{"a_b":{"c_d":"e"},"b_e":"x","e_f":[{"c_f":"e"}],"x_y":{"a_b":{"c_d":"e"},"b_e":"x","e_f":[{"c_f":"e"}]}}`)
	camelJson := []byte(`{"aB":{"cD":"e"},"bE":"x","eF":[{"cF":"e"}],"xY":{"aB":{"cD":"e"},"bE":"x","eF":[{"cF":"e"}]}}`)
	pascalJson := []byte(`{"AB":{"CD":"e"},"BE":"x","EF":[{"CF":"e"}],"XY":{"AB":{"CD":"e"},"BE":"x","EF":[{"CF":"e"}]}}`)

	testCases := []struct {
		name          string
		transformFunc func(json []byte) []byte
		input         []byte
		expected      []byte
	}{
		{
			name:          "camel-snake",
			transformFunc: ToSnakeJsonBytes,
			input:         camelJson,
			expected:      snakeJson,
		},
		{
			name:          "pascal-snake",
			transformFunc: ToSnakeJsonBytes,
			input:         pascalJson,
			expected:      snakeJson,
		},
		{
			name:          "snake-camel",
			transformFunc: ToCamelJsonBytes,
			input:         snakeJson,
			expected:      camelJson,
		},
		{
			name:          "pascal-camel",
			transformFunc: ToCamelJsonBytes,
			input:         pascalJson,
			expected:      camelJson,
		},
		{
			name:          "camel-pascal",
			transformFunc: ToPascalJsonBytes,
			input:         camelJson,
			expected:      pascalJson,
		},
		{
			name:          "snake-pascal",
			transformFunc: ToPascalJsonBytes,
			input:         snakeJson,
			expected:      pascalJson,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.transformFunc(tc.input)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
