package dataFormattingMonsters

import "strconv"



type Equipment struct {
	DropRate            float64        		`json:"drop_rate"`
	Power				int					`json:"power"`
	CriticalPlus		int					`json:"critical_plus"`
	AbilitySlots		AbilitySlots		`json:"ability_slots"`
	AttachedAbilities 	AttachedAbilities	`json:"attached_abilities"`
	WpnAbilities        []AutoAbility   	`json:"wpn_abilities"`
	ArmorAbilities      []AutoAbility 		`json:"armor_abilities"`
}


type AutoAbility struct {
	Ability		string	 `json:"ability"`
	Characters 	[]string `json:"characters"`
	Chance		float64	 `json:"chance"`
}


type AbilitySlots struct {
	MinAmount		int		`json:"min_amount"`
	MaxAmount		int		`json:"max_amount"`
	MinSlotsChance	float64	`json:"min_slots_chance"`
	MaxSlotsChance	float64	`json:"max_slots_chance"`
}


type AttachedAbilities struct {
	MinAmount 			int     `json:"min_amount"`
	MaxAmount 			int     `json:"max_amount"`
	MinAbilitiesChance	float64	`json:"min_abilities_chance"`
	MaxAbilitiesChance	float64	`json:"max_abilities_chance"`
}


type EquipmentOld struct {
	DropRate          float64           `json:"drop_rate"`
	SlotsAmount       string            `json:"slots_amount"`
	AttachedAbilities string            `json:"attached_abilities"`
	WpnAbilities      []WpnAbilityOld   `json:"wpn_abilities"`
	ArmorAbilities    []ArmorAbilityOld `json:"armor_abilities"`
}


type WpnAbilityOld struct {
	Ability    string `json:"ability"`
	Characters string `json:"characters,omitempty"`
}


type ArmorAbilityOld struct {
	Ability string `json:"ability"`
}





func (mon *MonsterOld) formatEquipment() Equipment {
	minSlots, maxSlots := mon.getMinMaxSlots()
	abilitySlots := AbilitySlots{
		MinAmount: minSlots,
		MaxAmount: maxSlots,
	}

	minAttached, maxAttached := mon.getMinMaxAttachedAbilities()
	attachedAbilities := AttachedAbilities{
		MinAmount: minAttached,
		MaxAmount: maxAttached,
	}


	return Equipment{
		DropRate:           mon.Equipment.DropRate,
		AbilitySlots: 		abilitySlots,
		AttachedAbilities: 	attachedAbilities,
		WpnAbilities:       mon.getWpnAbilities(),
		ArmorAbilities:     mon.getArmrAbilities(),
	}
}


func (mon *MonsterOld) getArmrAbilities() []AutoAbility {
	var armrAbilities []AutoAbility

	for _, armrAbilityOld := range mon.Equipment.ArmorAbilities {
		armrAbilities = append(armrAbilities, AutoAbility{
			Ability: armrAbilityOld.Ability,
			Characters: getCharacters(""),
		})
	}

	return armrAbilities
}

func (mon *MonsterOld) getWpnAbilities() []AutoAbility {
	var wpnAbilities []AutoAbility

	for _, wpnAbilityOld := range mon.Equipment.WpnAbilities {
		wpnAbilities = append(wpnAbilities, AutoAbility{
			Ability:    wpnAbilityOld.Ability,
			Characters: getCharacters(wpnAbilityOld.Characters),
		})
	}

	return wpnAbilities
}

func getCharacters(charString string) []string {
	var characters []string

	switch charString {
	case "Kimahri and Auron":
		characters = append(characters, []string{"kimahri", "auron"}...)
	case "Yuna and Lulu":
		characters = append(characters, []string{"yuna", "lulu"}...)
	case "except Yuna and Lulu":
		characters = append(characters, []string{"tidus", "wakka", "kimahri", "auron", "rikku"}...)
	default:
		characters = append(characters, "all")
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
