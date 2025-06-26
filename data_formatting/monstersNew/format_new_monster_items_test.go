package dataFormattingMonstersNew

import (
	"reflect"
	"testing"

	old "github.com/andreasSchauer/finalfantasyapi/data_formatting/monstersOld"
)

func TestGetItemData(t *testing.T) {
	tests := []struct {
		oldItem 		[]old.Item
		getSecondItem	bool
		expected 		*old.Item
	}{
		{
			oldItem: []old.Item{},
			getSecondItem: false,
			expected: nil,
		},
		{
			oldItem: []old.Item{
				{
					Item: "test name",
					Amount: 1,
				},
			},
			getSecondItem: false,
			expected: &old.Item{
				Item: "test name",
				Amount: 1,
			},
		},
		{
			oldItem: []old.Item{
				{
					Item: "test name",
					Amount: 1,
				},
				{
					Item: "second item",
					Amount: 3,
				},
			},
			getSecondItem: false,
			expected: &old.Item{
				Item: "test name",
				Amount: 1,
			},
		},
		{
			oldItem: []old.Item{
				{
					Item: "test name",
					Amount: 1,
				},
				{
					Item: "second item",
					Amount: 3,
				},
			},
			getSecondItem: true,
			expected: &old.Item{
				Item: "second item",
				Amount: 3,
			},
		},
	}
	
	for i, tc := range tests {
		actual := getItemData(tc.oldItem, tc.getSecondItem)
		if !reflect.DeepEqual(actual, tc.expected) {
			t.Errorf("getItemData: Testcase %d: input: %v, getSecondItem: %t. expected: %v, actual: %v", i, tc.oldItem, tc.getSecondItem, tc.expected, actual)
		}
	}
}