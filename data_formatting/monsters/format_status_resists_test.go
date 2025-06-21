package dataFormattingMonsters

import (
	"reflect"
	"testing"
)


func TestFormatResistances(t *testing.T) {
	tests := []struct {
		input		[]StatusResist
		expected	[]StatusResist
	}{
		{
			input: []StatusResist{
				{
					Status: "Silence",
					Resistance: 100,
				},
				{
					Status: "Dark",
					Resistance: 0,
				},
				{
					Status: "Sleep",
					Resistance: 101,
				},
			},
			expected: []StatusResist{
				{
					Status: "Silence",
					Resistance: 254,
				},
				{
					Status: "Dark",
					Resistance: 0,
				},
				{
					Status: "Sleep",
					Resistance: 101,
				},
			},
		},
	}

	for i, tc := range tests {
		actual := formatResistances(tc.input)
		if !reflect.DeepEqual(actual, tc.expected) {
			t.Errorf("formatResistances: Testcase %d: input: %v. expectedResistances: %v, actualResistances: %v", i, tc.input, tc.expected, actual)
		}
	}

}

func TestGetDemiResistance(t *testing.T) {
	tests := []struct {
		input 		float64
		expected 	int
	}{
		{
			input: float64(1),
			expected: 0,
		},
		{
			input: float64(0),
			expected: 100,
		},
	}
	
	for i, tc := range tests {
		elemResists := make(map[string]any)
		elemResists["gravity"] = tc.input

		mon := MonsterOld{
			ElemResists: elemResists,
		}
		actual := mon.getDemiResistance()
		if !reflect.DeepEqual(actual, tc.expected) {
			t.Errorf("getDemiResistance: Testcase %d: input: %v. expectedResistance: %d, actualResistance: %d", i, tc.input, tc.expected, actual)
		}
	}
}

func TestSplitDoom(t *testing.T) {
	tests := []struct {
		input 				[]int
		expectedResistance 	int
		expectedCountdown	*int
	}{
		{
			input: []int{0, 5},
			expectedResistance: 0,
			expectedCountdown: intpointer(5),
		},
		{
			input: []int{50, 5},
			expectedResistance: 50,
			expectedCountdown: intpointer(5),
		},
		{
			input: []int{100, 5},
			expectedResistance: 100,
			expectedCountdown: nil,
		},
		{
			input: []int{100, 0},
			expectedResistance: 100,
			expectedCountdown: nil,
		},
	}
	
	for i, tc := range tests {

		mon := MonsterOld{
			StatusResists: StatusResistsOld{
				Doom: tc.input,
			},
		}
		actualResistance, actualCountdown := mon.splitDoom()
		if !reflect.DeepEqual(actualResistance, tc.expectedResistance) {
			t.Errorf("splitDoom: Testcase %d: input: %v. expectedResistance: %d, actualResistance: %d", i, tc.input, tc.expectedResistance, actualResistance)
		}

		if !reflect.DeepEqual(actualCountdown, tc.expectedCountdown) {
			t.Errorf("splitDoom: Testcase %d: input: %v. expectedCountdown: %v, actualCountdown: %v", i, tc.input, tc.expectedCountdown, actualCountdown)
		}
	}
}


func TestSplitPoison(t *testing.T) {
	tests := []struct {
		input 				[]any
		expectedResistance 	int
		expectedRate		*float64
	}{
		{
			input: []any{float64(0), 0.5},
			expectedResistance: 0,
			expectedRate: floatpointer(0.5),
		},
		{
			input: []any{float64(0), float64(25)},
			expectedResistance: 0,
			expectedRate: floatpointer(25),
		},
		{
			input: []any{float64(100), float64(0)},
			expectedResistance: 100,
			expectedRate: nil,
		},
		{
			input: []any{float64(100), float64(25)},
			expectedResistance: 100,
			expectedRate: nil,
		},
	}
	
	for i, tc := range tests {

		mon := MonsterOld{
			StatusResists: StatusResistsOld{
				Poison: tc.input,
			},
		}
		actualResistance, actualRate := mon.splitPoison()
		if !reflect.DeepEqual(actualResistance, tc.expectedResistance) {
			t.Errorf("splitPoison: Testcase %d: input: %v. expectedResistance: %d, actualResistance: %d", i, tc.input, tc.expectedResistance, actualResistance)
		}

		if !reflect.DeepEqual(actualRate, tc.expectedRate) {
			t.Errorf("splitPoison: Testcase %d: input: %v. expectedRate: %v, actualRate: %v", i, tc.input, tc.expectedRate, actualRate)
		}
	}
}