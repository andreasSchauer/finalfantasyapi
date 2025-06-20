package dataFormatting

import (
	"reflect"
	"testing"
)

func TestFormatAllies(t *testing.T) {
	tests := []struct {
		allies 		[]string
		nameInput	string
		expected 	[]string
	}{
		{
			allies:		[]string{"string", "another string", "a third string"},
			nameInput:  "hello",
			expected:	[]string{"string", "another string", "a third string"},
		},
		{
			allies:		[]string{"string", "another string", "a third string"},
			nameInput:  "string",
			expected:	[]string{"another string", "a third string"},
		},
		{
			allies:		[]string{"string", "string", "a third string"},
			nameInput: 	"string",
			expected:	[]string{"a third string"},
		},
		{
			allies:		[]string{},
			nameInput:  "hello",
			expected:	[]string{},
		},
		
	}
	
	for i, tc := range tests {
		mon := MonsterOld{
			Allies: tc.allies,
		}
		actual := mon.formatAllies(tc.nameInput)
		if !reflect.DeepEqual(actual, tc.expected) {
			t.Errorf("formatAllies: Testcase %d: input: %s, expected: %v, actual: %v", i, tc.allies, tc.expected, actual)
		}
	}
}