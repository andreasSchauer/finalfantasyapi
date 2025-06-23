package dataFormattingMixes

import (
	"reflect"
	"testing"
)

func TestRejoinCombinations(t *testing.T) {
	tests := []struct {
		combinations []Combination
		expected     []string
	}{
		{
			combinations: []Combination{
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
			expected: []string{
				"potion + potion",
				"potion + hi-potion",
				"potion + al bhed potion",
				"hi-potion + x-potion",
				"elixir + phoenix down",
				"al bhed potion + al bhed potion",
				"bomb fragment + bomb core",
				"power sphere + speed sphere",
			},
		},
		{
			combinations: []Combination{},
			expected:     nil,
		},
	}

	for i, tc := range tests {
		actual := rejoinCombinations(tc.combinations)
		if !reflect.DeepEqual(actual, tc.expected) {
			t.Errorf("rejoinCombinations: Testcase %d: input: %v. expected: %v, actual: %v", i, tc.combinations, tc.expected, actual)
		}
	}
}
