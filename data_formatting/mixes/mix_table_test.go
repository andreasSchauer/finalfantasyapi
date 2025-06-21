package dataFormattingMixes

import (
	"reflect"
	"testing"
)

func TestRejoinCombinations(t *testing.T) {
	tests := []struct {
		combinations 	[]Combination
		expected 		[]string
	}{
		{
			combinations: []Combination{
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
			expected: nil,
		},
	}
	
	for i, tc := range tests {
		actual := rejoinCombinations(tc.combinations)
		if !reflect.DeepEqual(actual, tc.expected) {
			t.Errorf("rejoinCombinations: Testcase %d: input: %v. expected: %v, actual: %v", i, tc.combinations, tc.expected, actual)
		}
	}
}