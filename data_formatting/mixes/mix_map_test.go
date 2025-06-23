package dataFormattingMixes

import (
	"reflect"
	"testing"
)

func TestCreateMixMap(t *testing.T) {
	tests := []struct {
		mixName      string
		combinations []Combination
		expected     MixMap
	}{
		{
			mixName: "cool mix",
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
					Item2: "power sphere",
				},
				{
					Item1: "power sphere",
					Item2: "speed sphere",
				},
			},
			expected: MixMap{
				Name:               "cool mix",
				UniqueCombinations: 9,
				MixCombinations: map[string][]string{
					"potion":         {"potion", "hi-potion", "al bhed potion"},
					"hi-potion":      {"potion", "x-potion"},
					"x-potion":       {"hi-potion"},
					"elixir":         {"phoenix down"},
					"phoenix down":   {"elixir"},
					"al bhed potion": {"potion", "al bhed potion"},
					"bomb fragment":  {"bomb core"},
					"bomb core":      {"bomb fragment"},
					"power sphere":   {"power sphere", "speed sphere"},
					"speed sphere":   {"power sphere"},
				},
			},
		},
		{
			mixName:      "empty mix",
			combinations: []Combination{},
			expected: MixMap{
				Name:               "empty mix",
				UniqueCombinations: 0,
				MixCombinations:    make(map[string][]string),
			},
		},
	}

	for i, tc := range tests {
		actual := createMixMap(tc.mixName, tc.combinations)
		if !reflect.DeepEqual(actual, tc.expected) {
			t.Errorf("createMixMap: Testcase %d: input name: %s, input slice: %v. expected: %v, actual: %v", i, tc.mixName, tc.combinations, tc.expected, actual)
		}
	}
}

func TestGetJsonMixCombinations(t *testing.T) {
	tests := []struct {
		mixMap   MixMap
		expected []MixItems
	}{
		{
			mixMap: MixMap{
				Name:               "cool mix",
				UniqueCombinations: 8,
				MixCombinations: map[string][]string{
					"potion":         {"potion", "hi-potion", "al bhed potion"},
					"hi-potion":      {"potion", "x-potion"},
					"x-potion":       {"hi-potion"},
					"elixir":         {"phoenix down"},
					"phoenix down":   {"elixir"},
					"al bhed potion": {"potion", "al bhed potion"},
					"bomb fragment":  {"bomb core"},
					"bomb core":      {"bomb fragment"},
					"power sphere":   {"speed sphere"},
					"speed sphere":   {"power sphere"},
				},
			},
			expected: []MixItems{
				{
					FirstItem:           "potion",
					PossibleSecondItems: []string{"potion", "hi-potion", "al bhed potion"},
				},
				{
					FirstItem:           "hi-potion",
					PossibleSecondItems: []string{"potion", "x-potion"},
				},
				{
					FirstItem:           "x-potion",
					PossibleSecondItems: []string{"hi-potion"},
				},
				{
					FirstItem:           "elixir",
					PossibleSecondItems: []string{"phoenix down"},
				},
				{
					FirstItem:           "phoenix down",
					PossibleSecondItems: []string{"elixir"},
				},
				{
					FirstItem:           "al bhed potion",
					PossibleSecondItems: []string{"potion", "al bhed potion"},
				},
				{
					FirstItem:           "bomb fragment",
					PossibleSecondItems: []string{"bomb core"},
				},
				{
					FirstItem:           "bomb core",
					PossibleSecondItems: []string{"bomb fragment"},
				},
				{
					FirstItem:           "power sphere",
					PossibleSecondItems: []string{"speed sphere"},
				},
				{
					FirstItem:           "speed sphere",
					PossibleSecondItems: []string{"power sphere"},
				},
			},
		},
		{
			mixMap: MixMap{
				Name:               "cool mix",
				UniqueCombinations: 8,
				MixCombinations: map[string][]string{
					"potion":         {"potion", "hi-potion", "al bhed potion"},
					"x-potion":       {"hi-potion"},
					"elixir":         {"phoenix down"},
					"phoenix down":   {"elixir"},
					"al bhed potion": {"potion", "al bhed potion"},
					"hi-potion":      {"potion", "x-potion"},
					"bomb fragment":  {"bomb core"},
					"bomb core":      {"bomb fragment"},
					"power sphere":   {"speed sphere"},
					"speed sphere":   {"power sphere"},
				},
			},
			expected: []MixItems{
				{
					FirstItem:           "potion",
					PossibleSecondItems: []string{"potion", "hi-potion", "al bhed potion"},
				},
				{
					FirstItem:           "hi-potion",
					PossibleSecondItems: []string{"potion", "x-potion"},
				},
				{
					FirstItem:           "x-potion",
					PossibleSecondItems: []string{"hi-potion"},
				},
				{
					FirstItem:           "elixir",
					PossibleSecondItems: []string{"phoenix down"},
				},
				{
					FirstItem:           "phoenix down",
					PossibleSecondItems: []string{"elixir"},
				},
				{
					FirstItem:           "al bhed potion",
					PossibleSecondItems: []string{"potion", "al bhed potion"},
				},
				{
					FirstItem:           "bomb fragment",
					PossibleSecondItems: []string{"bomb core"},
				},
				{
					FirstItem:           "bomb core",
					PossibleSecondItems: []string{"bomb fragment"},
				},
				{
					FirstItem:           "power sphere",
					PossibleSecondItems: []string{"speed sphere"},
				},
				{
					FirstItem:           "speed sphere",
					PossibleSecondItems: []string{"power sphere"},
				},
			},
		},
		{
			mixMap: MixMap{
				Name:               "empty mix",
				UniqueCombinations: 0,
				MixCombinations:    make(map[string][]string),
			},
			expected: nil,
		},
	}

	for i, tc := range tests {
		mixMap := tc.mixMap
		actual := mixMap.getJsonMixCombinations()
		if !reflect.DeepEqual(actual, tc.expected) {
			t.Errorf("getJsonMixCombinations: Testcase %d: input: %v. expected: %v, actual: %v", i, tc.mixMap, tc.expected, actual)
		}
	}
}
