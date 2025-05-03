package dataFormatting

import (
	"reflect"
	"testing"
)

func TestSplitHP(t *testing.T) {
	tests := []struct {
		input 				[]any
		expectedHP 			int
		expectedOverkill	*int
	}{
		{
			input: []any{float64(1500), float64(2250)},
			expectedHP: 1500,
			expectedOverkill: intpointer(2250),
		},
		{
			input: []any{float64(1500), float64(1500)},
			expectedHP: 1500,
			expectedOverkill: intpointer(1500),
		},
		{
			input: []any{float64(0), float64(0)},
			expectedHP: 0,
			expectedOverkill: intpointer(0),
		},
		{
			input: []any{float64(1500), "-"},
			expectedHP: 1500,
			expectedOverkill: nil,
		},
		{
			input: []any{float64(1500), ""},
			expectedHP: 1500,
			expectedOverkill: nil,
		},
	}
	
	for i, tc := range tests {

		mon := MonsterOld{
			Stats: StatsOld{
				HP: tc.input,
			},
		}
		actualHP, actualOverkill := mon.splitHP()
		if !reflect.DeepEqual(actualHP, tc.expectedHP) {
			t.Errorf("splitHP: Testcase %d: input: %v. expectedHP: %d, actualHP: %d", i, tc.input, tc.expectedHP, actualHP)
		}

		if !reflect.DeepEqual(actualOverkill, tc.expectedOverkill) {
			t.Errorf("splitHP: Testcase %d: input: %v. expectedOverkill: %v, actualOverkill: %v", i, tc.input, tc.expectedOverkill, actualOverkill)
		}
	}
}


func intpointer(i int) *int {
	return &i
}

func floatpointer(f float64) *float64 {
	return &f
}