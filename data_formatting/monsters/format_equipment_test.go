package dataFormattingMonsters

import (
	"reflect"
	"testing"
)

func TestGetMinMaxFuncs(t *testing.T) {
	tests := []struct {
		input       string
		expectedMin int
		expectedMax int
	}{
		{
			input:       "1-3",
			expectedMin: 1,
			expectedMax: 3,
		},
		{
			input:       "0",
			expectedMin: 0,
			expectedMax: 0,
		},
		{
			input:       "-",
			expectedMin: 0,
			expectedMax: 0,
		},
		{
			input:       "",
			expectedMin: 0,
			expectedMax: 0,
		},
	}

	for i, tc := range tests {
		mon := MonsterOld{
			Equipment: EquipmentOld{
				SlotsAmount:       tc.input,
				AttachedAbilities: tc.input,
			},
		}
		minSlots, maxSlots := mon.getMinMaxSlots()
		minAttached, maxAttached := mon.getMinMaxAttachedAbilities()

		if !reflect.DeepEqual(minSlots, tc.expectedMin) {
			t.Errorf("getMinMaxSlots: Testcase %d: input: %s. expectedMin: %d, actualMin: %d", i, tc.input, tc.expectedMin, minSlots)
		}

		if !reflect.DeepEqual(maxSlots, tc.expectedMax) {
			t.Errorf("getMinMaxSlots: Testcase %d: input: %s. expectedMax: %d, actualMax: %d", i, tc.input, tc.expectedMax, maxSlots)
		}

		if !reflect.DeepEqual(minAttached, tc.expectedMin) {
			t.Errorf("getMinMaxAttachedAbilities: Testcase %d: input: %s. expectedMin: %d, actualMin: %d", i, tc.input, tc.expectedMin, minAttached)
		}

		if !reflect.DeepEqual(maxAttached, tc.expectedMax) {
			t.Errorf("getMinMaxAttachedAbilities: Testcase %d: input: %s. expectedMax: %d, actualMax: %d", i, tc.input, tc.expectedMax, maxAttached)
		}
	}
}

func TestGetWpnAbilities(t *testing.T) {
	tests := []struct {
		input    []WpnAbilityOld
		expected []WpnAbility
	}{
		{
			input: []WpnAbilityOld{
				{
					Ability: "test-ability",
					Characters: "Kimahri and Auron",
				},
				{
					Ability: "test-ability-2",
					Characters: "Yuna and Lulu",
				},
				{
					Ability: "test-ability-3",
					Characters: "Wakka and Auron",
				},
				{
					Ability: "test-ability-4",
					Characters: "except Yuna and Lulu",
				},
				{
					Ability: "test-ability-5",
				},
			},
			expected: []WpnAbility{
				{
					Ability: "test-ability",
					Characters: Characters{
						KimahriAndAuron: true,
					},
				},
				{
					Ability: "test-ability-2",
					Characters: Characters{
						YunaAndLulu: true,
					},
				},
				{
					Ability: "test-ability-3",
					Characters: Characters{},
				},
				{
					Ability: "test-ability-4",
					Characters: Characters{
						ExceptYunaAndLulu: true,
					},
				},
				{
					Ability: "test-ability-5",
				},
			},
		},
		{
			input: nil,
			expected: nil,
		},
	}

	for i, tc := range tests {
		mon := MonsterOld{
			Equipment: EquipmentOld{
				WpnAbilities: tc.input,
			},
		}
		actual := mon.getWpnAbilities()

		if !reflect.DeepEqual(actual, tc.expected) {
			t.Errorf("getWpnAbilities: Testcase %d: input: %v. expected: %v, actual: %v", i, tc.input, tc.expected, actual)
		}
	}
}
