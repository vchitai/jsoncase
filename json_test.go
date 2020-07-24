package jsoncase

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTransformVal(t *testing.T) {
	snakeVal := map[string]interface{}{
		"a_b": map[string]interface{}{
			"c_d": "e",
		},
		"b_e": "x",
		"e_f": []interface{}{
			map[string]interface{}{
				"c_f": "e",
			},
		},
		"x_y": map[string]interface{}{
			"a_b": map[string]interface{}{
				"c_d": "e",
			},
			"b_e": "x",
			"e_f": []map[string]interface{}{
				{
					"c_f": "e",
				},
			},
		},
	}
	camelVal := map[string]interface{}{
		"aB": map[string]interface{}{
			"cD": "e",
		},
		"bE": "x",
		"eF": []interface{}{
			map[string]interface{}{
				"cF": "e",
			},
		},
		"xY": map[string]interface{}{
			"aB": map[string]interface{}{
				"cD": "e",
			},
			"bE": "x",
			"eF": []map[string]interface{}{
				{
					"cF": "e",
				},
			},
		},
	}
	pascalVal := map[string]interface{}{
		"AB": map[string]interface{}{
			"CD": "e",
		},
		"BE": "x",
		"EF": []interface{}{
			map[string]interface{}{
				"CF": "e",
			},
		},
		"XY": map[string]interface{}{
			"AB": map[string]interface{}{
				"CD": "e",
			},
			"BE": "x",
			"EF": []map[string]interface{}{
				{
					"CF": "e",
				},
			},
		},
	}

	testCases := []struct {
		name          string
		transformFunc func(json interface{}) interface{}
		input         interface{}
		expected      interface{}
	}{
		{
			name:          "camel-snake",
			transformFunc: ToSnakeVal,
			input:         camelVal,
			expected:      snakeVal,
		},
		{
			name:          "pascal-snake",
			transformFunc: ToSnakeVal,
			input:         pascalVal,
			expected:      snakeVal,
		},
		{
			name:          "snake-camel",
			transformFunc: ToCamelVal,
			input:         snakeVal,
			expected:      camelVal,
		},
		{
			name:          "pascal-camel",
			transformFunc: ToCamelVal,
			input:         pascalVal,
			expected:      camelVal,
		},
		{
			name:          "camel-pascal",
			transformFunc: ToPascalVal,
			input:         camelVal,
			expected:      pascalVal,
		},
		{
			name:          "snake-pascal",
			transformFunc: ToPascalVal,
			input:         snakeVal,
			expected:      pascalVal,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.transformFunc(tc.input)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
