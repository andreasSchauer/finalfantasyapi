package dataFormattingMonstersNew

import(
	old "github.com/andreasSchauer/finalfantasyapi/data_formatting/monstersOld"
)

type StatusResistsNew struct {
	Immunities		[]string			`json:"immunities"`
	NegResists      []old.StatusResist 	`json:"neg_resists"`
	PosResists      []old.StatusResist 	`json:"pos_resists"`
	MiscResists     []old.StatusResist 	`json:"misc_resists"`
}


func formatStatusResists(StatusResists old.StatusResists) *StatusResistsNew {
	var immunities []string
	negResists := populateImmunities(&immunities, StatusResists.NegResists)
	posResists := populateImmunities(&immunities, StatusResists.PosResists)
	miscResists := populateImmunities(&immunities, StatusResists.MiscResists)

	return &StatusResistsNew{
		Immunities: immunities,
		NegResists: negResists,
		PosResists: posResists,
		MiscResists: miscResists,
	}
}


func populateImmunities(immunities *[]string, resistances []old.StatusResist) []old.StatusResist {
	var resistancesFiltered []old.StatusResist
	
	for _, condition := range resistances {
		if condition.Resistance == 254 {
			*immunities = append(*immunities, condition.Status)
			continue
		}

		resistancesFiltered = append(resistancesFiltered, condition)
	}

	return resistancesFiltered
}