package dataFormatting

import "strconv"

func (mon *MonsterOld) formatEquipment() Equipment {
	minSlots, maxSlots := mon.getMinMaxSlots()
	minAttached, maxAttached := mon.getMinMaxAttachedAbilities()

	return Equipment{
		DropRate:             mon.Equipment.DropRate,
		MinSlotsAmount:       minSlots,
		MaxSlotsAmount:       maxSlots,
		MinAttachedAbilities: minAttached,
		MaxAttachedAbilities: maxAttached,
		WpnAbilities:         mon.getWpnAbilities(),
		ArmorAbilities:       mon.getArmrAbilities(),
	}
}

func (mon *MonsterOld) getArmrAbilities() []ArmorAbility {
	var armrAbilities []ArmorAbility

	for _, armrAbilityOld := range mon.Equipment.ArmorAbilities {
		armrAbilities = append(armrAbilities, ArmorAbility(armrAbilityOld))
	}

	return armrAbilities
}

func (mon *MonsterOld) getWpnAbilities() []WpnAbility {
	var wpnAbilities []WpnAbility

	for _, wpnAbilityOld := range mon.Equipment.WpnAbilities {
		wpnAbility := WpnAbility{
			Ability:    wpnAbilityOld.Ability,
			Characters: getCharacters(wpnAbilityOld.Characters),
		}
		wpnAbilities = append(wpnAbilities, wpnAbility)
	}

	return wpnAbilities
}

func getCharacters(charString string) Characters {
	characters := Characters{}

	switch charString {
	case "Kimahri and Auron":
		characters.KimahriAndAuron = true
	case "Yuna and Lulu":
		characters.YunaAndLulu = true
	case "except Yuna and Lulu":
		characters.ExceptYunaAndLulu = true
	}

	return characters
}

func (mon *MonsterOld) getMinMaxSlots() (int, int) {
	slotsAmount := mon.Equipment.SlotsAmount

	if len(slotsAmount) == 0 {
		return 0, 0
	}

	minSlots, _ := strconv.Atoi(string(slotsAmount[0]))
	maxSlots, _ := strconv.Atoi(string(slotsAmount[len(slotsAmount)-1]))

	return minSlots, maxSlots
}

func (mon *MonsterOld) getMinMaxAttachedAbilities() (int, int) {
	attachedAbilities := mon.Equipment.AttachedAbilities

	if len(attachedAbilities) == 0 {
		return 0, 0
	}

	minAttached, _ := strconv.Atoi(string(attachedAbilities[0]))
	maxAttached, _ := strconv.Atoi(string(attachedAbilities[len(attachedAbilities)-1]))

	return minAttached, maxAttached
}
