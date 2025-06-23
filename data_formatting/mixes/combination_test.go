package dataFormattingMixes

import (
	"reflect"
	"testing"
)

func TestGetCombinations(t *testing.T) {
	tests := []struct {
		combinationStrings []string
		expected           []Combination
	}{
		{
			combinationStrings: []string{
				"al bhed potion + al bhed potion",
				"al bhed potion + potion",
				"elixir + phoenix down",
				"power sphere + speed sphere",
				"",
				"bomb core + bomb fragment",
				"hi-potion + x-potion",
				"",
				"hi-potion + potion",
				"potion + potion",
			},
			expected: []Combination{
				{
					Item1: "potion",
					Item2: "potion",
				},
				{
					Item1: "potion",
					Item2: "hi-potion",
				},
				{
					Item1: "potion",
					Item2: "al bhed potion",
				},
				{
					Item1: "hi-potion",
					Item2: "x-potion",
				},
				{
					Item1: "elixir",
					Item2: "phoenix down",
				},
				{
					Item1: "al bhed potion",
					Item2: "al bhed potion",
				},
				{
					Item1: "bomb fragment",
					Item2: "bomb core",
				},
				{
					Item1: "power sphere",
					Item2: "speed sphere",
				},
			},
		},
		{
			combinationStrings: []string{},
			expected:           nil,
		},
	}

	for i, tc := range tests {
		actual := getCombinations(tc.combinationStrings)
		if !reflect.DeepEqual(actual, tc.expected) {
			t.Errorf("getCombinations: Testcase %d: input: %v. expected: %v, actual: %v", i, tc.combinationStrings, tc.expected, actual)
		}
	}
}
