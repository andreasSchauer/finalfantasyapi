package dataFormatting

import (
	"reflect"
	"sort"
	"testing"
)

func TestFormatElemResists(t *testing.T) {
	tests := []struct {
		elemResists map[string]any
		expected 	[]ElemResist
	}{
		{
			elemResists: map[string]any{
				"fire": 		float64(1.5),
				"lightning": 	float64(1),
				"water": 		float64(0.5),
				"ice": 			float64(0),
				"holy": 		float64(-1),
				"gravity": 		float64(1),
			},
			expected:		[]ElemResist{
				{
					Element: "fire",
					Affinity: "weak",
				},
				{
					Element: "lightning",
					Affinity: "neutral",
				},
				{
					Element: "water",
					Affinity: "halved",
				},
				{
					Element: "ice",
					Affinity: "immune",
				},
				{
					Element: "holy",
					Affinity: "absorb",
				},
			},
		},
		{
			elemResists: map[string]any{
				"fire": 	"varies",
				"lightning": "",
			},
			expected:		[]ElemResist{
				{
					Element: "fire",
					Affinity: "varies",
				},
				{
					Element: "lightning",
					Affinity: "",
				},
			},
		},
		{
			elemResists: map[string]any{
				"fire": 	 nil,
				"lightning": "paul",
			},
			expected:		[]ElemResist{
				{
					Element: "fire",
					Affinity: "",
				},
				{
					Element: "lightning",
					Affinity: "",
				},
			},
		},
		{
			elemResists: map[string]any{
				"fire": 	 float64(2),
				"lightning": float64(1.5),
			},
			expected:		[]ElemResist{
				{
					Element: "fire",
					Affinity: "",
				},
				{
					Element: "lightning",
					Affinity: "weak",
				},
			},
		},
	}
	
	for i, tc := range tests {
		mon := MonsterOld{
			ElemResists: tc.elemResists,
		}
		actual := mon.formatElemResist()
		sort.SliceStable(actual, func(i, j int) bool { return actual[i].Element < actual[j].Element })
		sort.SliceStable(tc.expected, func(i, j int) bool { return tc.expected[i].Element < tc.expected[j].Element })
		if !reflect.DeepEqual(actual, tc.expected) {
			t.Errorf("formatElemResist: Testcase %d: input: %v. expected: %v, actual: %v", i, tc.elemResists, tc.expected, actual)
		}
	}
}