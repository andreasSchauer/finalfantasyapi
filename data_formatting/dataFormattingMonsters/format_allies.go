package dataFormattingMonsters


func (mon *MonsterOld) formatAllies(name string) []string {
	alliesFormatted := make([]string, 0)

	for _, mon := range(mon.Allies) {
		if mon != name {
			alliesFormatted = append(alliesFormatted, mon)
		}
	}

	return alliesFormatted
}