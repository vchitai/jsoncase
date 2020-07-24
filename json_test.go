package jsoncase

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTransformVal(t *testing.T) {
	snakeMap := map[string]interface{}{
		"a_b": map[string]interface{}{
			"c_d": "e",
		},
		"b_e": "x",
		"e_f": []interface{}{
			map[string]interface{}{
				"c_f": "e",
			},
			"e",
			1,
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
	snakeSlice := []map[string]interface{}{
		{
			"a_b": map[string]interface{}{
				"c_d": "e",
			},
			"b_e": "x",
			"e_f": []interface{}{
				map[string]interface{}{
					"c_f": "e",
				},
				"e",
				1,
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
		},
	}
	camelMap := map[string]interface{}{
		"aB": map[string]interface{}{
			"cD": "e",
		},
		"bE": "x",
		"eF": []interface{}{
			map[string]interface{}{
				"cF": "e",
			},
			"e",
			1,
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
	camelSlice := []map[string]interface{}{
		{
			"aB": map[string]interface{}{
				"cD": "e",
			},
			"bE": "x",
			"eF": []interface{}{
				map[string]interface{}{
					"cF": "e",
				},
				"e",
				1,
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
		},
	}
	pascalMap := map[string]interface{}{
		"AB": map[string]interface{}{
			"CD": "e",
		},
		"BE": "x",
		"EF": []interface{}{
			map[string]interface{}{
				"CF": "e",
			},
			"e",
			1,
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
	pascalSlice := []map[string]interface{}{
		{
			"AB": map[string]interface{}{
				"CD": "e",
			},
			"BE": "x",
			"EF": []interface{}{
				map[string]interface{}{
					"CF": "e",
				},
				"e",
				1,
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
		},
	}

	testCases := []struct {
		name          string
		transformFunc func(json interface{}) interface{}
		input         interface{}
		expected      interface{}
	}{
		{
			name:          "camel-snake map",
			transformFunc: ToSnakeVal,
			input:         camelMap,
			expected:      snakeMap,
		},
		{
			name:          "pascal-snake map",
			transformFunc: ToSnakeVal,
			input:         pascalMap,
			expected:      snakeMap,
		},
		{
			name:          "snake-camel map",
			transformFunc: ToCamelVal,
			input:         snakeMap,
			expected:      camelMap,
		},
		{
			name:          "pascal-camel map",
			transformFunc: ToCamelVal,
			input:         pascalMap,
			expected:      camelMap,
		},
		{
			name:          "camel-pascal map",
			transformFunc: ToPascalVal,
			input:         camelMap,
			expected:      pascalMap,
		},
		{
			name:          "snake-pascal map",
			transformFunc: ToPascalVal,
			input:         snakeMap,
			expected:      pascalMap,
		},
		{
			name:          "camel-snake slice",
			transformFunc: ToSnakeVal,
			input:         camelSlice,
			expected:      snakeSlice,
		},
		{
			name:          "pascal-snake slice",
			transformFunc: ToSnakeVal,
			input:         pascalSlice,
			expected:      snakeSlice,
		},
		{
			name:          "snake-camel slice",
			transformFunc: ToCamelVal,
			input:         snakeSlice,
			expected:      camelSlice,
		},
		{
			name:          "pascal-camel slice",
			transformFunc: ToCamelVal,
			input:         pascalSlice,
			expected:      camelSlice,
		},
		{
			name:          "camel-pascal slice",
			transformFunc: ToPascalVal,
			input:         camelSlice,
			expected:      pascalSlice,
		},
		{
			name:          "snake-pascal slice",
			transformFunc: ToPascalVal,
			input:         snakeSlice,
			expected:      pascalSlice,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.transformFunc(tc.input)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestTransformMap(t *testing.T) {
	snakeMap := map[string]interface{}{
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
	camelMap := map[string]interface{}{
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
	pascalMap := map[string]interface{}{
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
		transformFunc func(json map[string]interface{}) map[string]interface{}
		input         map[string]interface{}
		expected      interface{}
	}{
		{
			name:          "camel-snake",
			transformFunc: TransformMap(toSnakeString),
			input:         camelMap,
			expected:      snakeMap,
		},
		{
			name:          "pascal-snake",
			transformFunc: TransformMap(toSnakeString),
			input:         pascalMap,
			expected:      snakeMap,
		},
		{
			name:          "snake-camel",
			transformFunc: TransformMap(toCamelString),
			input:         snakeMap,
			expected:      camelMap,
		},
		{
			name:          "pascal-camel",
			transformFunc: TransformMap(toCamelString),
			input:         pascalMap,
			expected:      camelMap,
		},
		{
			name:          "camel-pascal",
			transformFunc: TransformMap(toPascalString),
			input:         camelMap,
			expected:      pascalMap,
		},
		{
			name:          "snake-pascal",
			transformFunc: TransformMap(toPascalString),
			input:         snakeMap,
			expected:      pascalMap,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.transformFunc(tc.input)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
