package dataFormattingMonsters



type Stat struct {
	Name		string  `json:"name"`
	Value		int		`json:"value"`
}


type StatsOld struct {
	HP         []any `json:"hp"`
	MP         int   `json:"mp"`
	Strength   int   `json:"strength"`
	Defence    int   `json:"defence"`
	Magic      int   `json:"magic"`
	MagDefence int   `json:"mag_defence"`
	Agility    int   `json:"agility"`
	Luck       int   `json:"luck"`
	Evasion    int   `json:"evasion"`
	Accuracy   int   `json:"accuracy"`
}



func (mon *MonsterOld) formatStats() ([]Stat, *int) {
	statsOld := mon.Stats
	hp, overkillDamage := mon.splitHP()

	return []Stat{
		{
			Name: 	"hp",
			Value:	hp,
		},
		{
			Name: "mp",
			Value: statsOld.MP,
		},
		{
			Name: "strength",
			Value: statsOld.Strength,
		},
		{
			Name: "defense",
			Value: statsOld.Defence,
		},
		{
			Name: "magic",
			Value: statsOld.Magic,
		},
		{
			Name: "magic defense",
			Value: statsOld.MagDefence,
		},
		{
			Name: "agility",
			Value: statsOld.Agility,
		},
		{
			Name: "luck",
			Value: statsOld.Luck,
		},
		{
			Name: "evasion",
			Value: statsOld.Evasion,
		},
		{
			Name: "accuracy",
			Value: statsOld.Accuracy,
		},
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
