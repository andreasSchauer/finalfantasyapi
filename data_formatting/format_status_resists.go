package dataFormatting

func (mon *MonsterOld) formatStatusResists() StatusResists {
	poisonResist, poisonRate := mon.splitPoison()
	doomResist, doomCountdown := mon.splitDoom()
	
	return StatusResists{
		PoisonRate: poisonRate,
		DoomCountdown: doomCountdown,
		ZanmatoLevel: mon.StatusResists.ZanmatoLevel,
		IsImmuneToLife: mon.isImmuneToLife(),
		Resistances: mon.getStatusResistances(poisonResist, doomResist),
	}
}


func (mon *MonsterOld) getStatusResistances(poisonResist, doomResist int) []StatusResist {
	return []StatusResist{
		{
			Status: "Silence",
			Resistance: mon.StatusResists.Silence,
		},
		{
			Status: "Sleep",
			Resistance: mon.StatusResists.Sleep,
		},
		{
			Status: "Dark",
			Resistance: mon.StatusResists.Dark,
		},
		{
			Status: "Poison",
			Resistance: poisonResist,
		},
		{
			Status: "Petrify",
			Resistance: mon.StatusResists.Petrify,
		},
		{
			Status: "Slow",
			Resistance: mon.StatusResists.Slow,
		},
		{
			Status: "Zombie",
			Resistance: mon.StatusResists.Zombie,
		},
		{
			Status: "Power Break",
			Resistance: mon.StatusResists.PowerBreak,
		},
		{
			Status: "Magic Break",
			Resistance: mon.StatusResists.MagicBreak,
		},
		{
			Status: "Armour Break",
			Resistance: mon.StatusResists.ArmourBreak,
		},
		{
			Status: "Mental Break",
			Resistance: mon.StatusResists.MentalBreak,
		},
		{
			Status: "Threaten",
			Resistance: mon.StatusResists.Threaten,
		},
		{
			Status: "Death",
			Resistance: mon.StatusResists.Death,
		},
		{
			Status: "Provoke",
			Resistance: mon.StatusResists.Provoke,
		},
		{
			Status: "Doom",
			Resistance: doomResist,
		},
		{
			Status: "Delay",
			Resistance: mon.StatusResists.Delay,
		},
		{
			Status: "Eject",
			Resistance: mon.StatusResists.Eject,
		},
	}
}


func (mon *MonsterOld) isImmuneToLife() bool {
	return mon.StatusResists.Life == 100
}


func (mon *MonsterOld) splitPoison() (int, *float64) {
	poisonData := mon.StatusResists.Poison
	poisonResist := int(poisonData[0].(float64))
	poisonRateVal := poisonData[1].(float64)
	var poisonRate *float64

	if poisonResist < 100 {
		poisonRate = &poisonRateVal
	}

	return poisonResist, poisonRate
}


func (mon *MonsterOld) splitDoom() (int, *int) {
	doomData := mon.StatusResists.Doom
	doomResist := doomData[0]
	doomCountdownVal := doomData[1]
	var doomCountdown *int

	if doomResist < 100 {
		doomCountdown = &doomCountdownVal
	}

	return doomResist, doomCountdown
}

/*

func (mon *MonsterOld) splitHP() (int, *int) {
	hpData := mon.Stats.HP
	hp := int(hpData[0].(float64))
	overkillDamageData := hpData[1]
	var overkillDamage *int

	if val, ok := overkillDamageData.(float64); ok {
		v := int(val)
		overkillDamage = &v
	}

	return hp, overkillDamage
}

*/
