package dataFormatting

func (mon *MonsterOld) formatStats() (Stats, *int) {
	statsOld := mon.Stats
	hp, overkillDamage := mon.splitHP()

	return Stats{
		HP:         hp,
		MP:         statsOld.MP,
		Strength:   statsOld.Strength,
		Defense:    statsOld.Defence,
		Magic:      statsOld.Magic,
		MagDefense: statsOld.MagDefence,
		Agility:    statsOld.Agility,
		Luck:       statsOld.Luck,
		Evasion:    statsOld.Evasion,
		Accuracy:   statsOld.Accuracy,
	}, overkillDamage
}

func (mon *MonsterOld) splitHP() (int, *int) {
	hpData := mon.Stats.HP
	hp := int(hpData[0].(float64))
	overkillDamageVal := hpData[1]
	var overkillDamage *int

	if val, ok := overkillDamageVal.(float64); ok {
		v := int(val)
		overkillDamage = &v
	}

	return hp, overkillDamage
}
