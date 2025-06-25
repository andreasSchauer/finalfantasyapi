package dataFormattingMonsters

import (
	"reflect"
	"testing"
)

func TestGetItems(t *testing.T) {
	tests := []struct {
		itemData 	[]any
		expected 	[]Item
	}{
		{
			itemData: []any{"cool item name", float64(2)},
			expected: []Item{
				{
					Item: "cool item name",
					Amount: 2,
				},
			},
		},
		{
			itemData: []any{[]any{"cool item name", float64(2)}, []any{"even cooler item", float64(5)}},
			expected: []Item{
				{
					Item: "cool item name",
					Amount: 2,
				},
				{
					Item: "even cooler item",
					Amount: 5,
				},
			},
		},
		{
			itemData: nil,
			expected: []Item{},
		},
	}
	
	for i, tc := range tests {
		key := "drop_normal"
		items := make(map[string][]any)
		items[key] = tc.itemData

		mon := MonsterOld{
			Items: items,
		}
		actual := mon.getItems(key)
		if !reflect.DeepEqual(actual, tc.expected) {
			t.Errorf("getItems: Testcase %d: input: %v. expected: %v, actual: %v", i, tc.itemData, tc.expected, actual)
		}
	}
}