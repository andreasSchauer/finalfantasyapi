package dataFormattingMixes

import (
	"reflect"
	"testing"
)

func TestGetCombinations(t *testing.T) {
	tests := []struct {
		combinationStrings 	[]string
		expected     		[]Combination
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
					item1: "potion",
					item2: "potion",
				},
				{
					item1: "potion",
					item2: "hi-potion",
				},
				{
					item1: "potion",
					item2: "al bhed potion",
				},
				{
					item1: "hi-potion",
					item2: "x-potion",
				},
				{
					item1: "elixir",
					item2: "phoenix down",
				},
				{
					item1: "al bhed potion",
					item2: "al bhed potion",
				},
				{
					item1: "bomb fragment",
					item2: "bomb core",
				},
				{
					item1: "power sphere",
					item2: "speed sphere",
				},
			},
		},
		{
			combinationStrings: []string{},
			expected:     nil,
		},
	}

	for i, tc := range tests {
		actual := getCombinations(tc.combinationStrings)
		if !reflect.DeepEqual(actual, tc.expected) {
			t.Errorf("getCombinations: Testcase %d: input: %v. expected: %v, actual: %v", i, tc.combinationStrings, tc.expected, actual)
		}
	}
}
