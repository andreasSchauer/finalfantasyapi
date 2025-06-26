package dataFormattingMonstersNew

import(
	old "github.com/andreasSchauer/finalfantasyapi/data_formatting/monstersOld"
)


type AlteredStateOld struct {
	Condition		string				`json:"condition"`
	Properties		[]string			`json:"properties"`
	Position		string				`json:"position"`
	Stats           []old.Stat        	`json:"stats"`
	ElemResists     []old.ElemResist  	`json:"elem_resists"`
	StatusResists   *old.StatusResists 	`json:"status_resists"`
}



type AlteredStateNew struct {
	Condition		string				`json:"condition"`
	Properties		[]string			`json:"properties"`
	Position		string				`json:"position"`
	Stats           []old.Stat        	`json:"stats"`
	ElemResists     []old.ElemResist  	`json:"elem_resists"`
	StatusResists   *StatusResistsNew 	`json:"status_resists"`
}



func (mon *MonsterOld) getAlteredStates() []AlteredStateNew {
	alteredStates := []AlteredStateNew{}

	if len(mon.AlteredStates) == 0 {
		return alteredStates
	}

	for _, state := range mon.AlteredStates {
		var statusResists *StatusResistsNew
		if state.StatusResists == nil {
			statusResists = nil
		} else {
			statusResists = formatStatusResists(*state.StatusResists)
		}

		alteredStates = append(alteredStates, AlteredStateNew{
			Condition: state.Condition,
			Properties: state.Properties,
			Position: state.Position,
			Stats: state.Stats,
			ElemResists: state.ElemResists,
			StatusResists: statusResists,
		})
	}

	return alteredStates
}