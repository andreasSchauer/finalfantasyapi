package dataFormatting

import (

)


func (mon *MonsterOld) formatElemResist() []ElemResist {
	elemResists := []ElemResist{
		mon.createElemResist("fire"),
		mon.createElemResist("lightning"),
		mon.createElemResist("water"),
		mon.createElemResist("ice"),
		mon.createElemResist("holy"),
	}


	return elemResists
}


func (mon *MonsterOld) createElemResist(element string) ElemResist {
	return ElemResist{
		Element: element,
		Affinity: convertAffinity(mon.ElemResists[element]),
	}
}

func convertAffinity(multiplier any) string {
	if val, ok := multiplier.(string); ok && val == "varies" {
		return val
	}
	
	if val, ok := multiplier.(float64); ok {
		switch val {
		case 1.5:
			return "weak"
		case 1:
			return "neutral"
		case 0.5:
			return "halved"
		case 0:
			return "immune"
		case -1:
			return "absorb"
		}
	}
	
	return ""
}