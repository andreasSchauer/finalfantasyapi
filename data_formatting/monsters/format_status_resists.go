package dataFormattingMonsters

type StatusResistsData struct {
	StatusResists   StatusResists
	PoisonRate      *float64       `json:"poison_rate"`
	DoomCountdown   *int           `json:"doom_countdown"`
	ThreatenCounter int            `json:"threaten_counter"`
	ZanmatoLevel    int            `json:"zanmato_level"`
}

type StatusResists struct {
	NegResists      []StatusResist `json:"neg_resists"`
	PosResists      []StatusResist `json:"pos_resists"`
	MiscResists     []StatusResist `json:"misc_resists"`
}

type StatusResist struct {
	Status     string `json:"status"`
	Resistance int    `json:"resistance"`
}

type StatusResistsOld struct {
	Silence      int   `json:"silence"`
	Sleep        int   `json:"sleep"`
	Dark         int   `json:"dark"`
	Poison       []any `json:"poison"`
	Petrify      int   `json:"petrify"`
	Slow         int   `json:"slow"`
	Zombie       int   `json:"zombie"`
	Life         int   `json:"life"`
	PowerBreak   int   `json:"power_break"`
	MagicBreak   int   `json:"magic_break"`
	ArmorBreak   int   `json:"armor_break"`
	MentalBreak  int   `json:"mental_break"`
	Threaten     int   `json:"threaten"`
	Death        int   `json:"death"`
	Provoke      int   `json:"provoke"`
	Doom         []int `json:"doom"`
	NulSpells    int   `json:"nul_spells"`
	Shell        int   `json:"shell"`
	Protect      int   `json:"protect"`
	Reflect      int   `json:"reflect"`
	Haste        int   `json:"haste"`
	Regen        int   `json:"regen"`
	Distiller    int   `json:"distiller"`
	Sensor       int   `json:"sensor"`
	Scan         int   `json:"scan"`
	Delay        int   `json:"delay"`
	Eject        int   `json:"eject"`
	Berserk      int   `json:"berserk"`
	ZanmatoLevel int   `json:"zanmato_level"`
}

func (mon *MonsterOld) formatStatusResists() StatusResistsData {
	poisonResist, poisonRate := mon.splitPoison()
	doomResist, doomCountdown := mon.splitDoom()
	negResists := mon.getNegResists(poisonResist, doomResist)
	posResists := mon.getPosResists()
	miscResists := mon.getMiscResists()

	return StatusResistsData{
		PoisonRate:    poisonRate,
		DoomCountdown: doomCountdown,
		ZanmatoLevel:  mon.StatusResists.ZanmatoLevel,
		StatusResists: StatusResists{
			NegResists:    formatResistances(negResists),
			PosResists:    formatResistances(posResists),
			MiscResists:   formatResistances(miscResists),
		},
	}
}

func formatResistances(resistances []StatusResist) []StatusResist {
	resistances_formatted := []StatusResist{}

	for i := range resistances {
		resistances_formatted = append(resistances_formatted, resistances[i])
		if resistances_formatted[i].Resistance == 100 {
			resistances_formatted[i].Resistance = 254
		}
	}

	return resistances_formatted
}

func (mon *MonsterOld) getNegResists(poisonResist, doomResist int) []StatusResist {
	return filterStatusResists([]StatusResist{
		{
			Status:     "berserk",
			Resistance: mon.StatusResists.Berserk,
		},
		{
			Status:     "confuse",
			Resistance: 0,
		},
		{
			Status:     "dark",
			Resistance: mon.StatusResists.Dark,
		},
		{
			Status:     "death",
			Resistance: mon.StatusResists.Death,
		},
		{
			Status:     "delay",
			Resistance: mon.StatusResists.Delay,
		},
		{
			Status:     "doom",
			Resistance: doomResist,
		},
		{
			Status:     "eject",
			Resistance: mon.StatusResists.Eject,
		},
		{
			Status:     "petrify",
			Resistance: mon.StatusResists.Petrify,
		},
		{
			Status:     "poison",
			Resistance: poisonResist,
		},
		{
			Status:     "provoke",
			Resistance: mon.StatusResists.Provoke,
		},
		{
			Status:     "silence",
			Resistance: mon.StatusResists.Silence,
		},
		{
			Status:     "sleep",
			Resistance: mon.StatusResists.Sleep,
		},
		{
			Status:     "slow",
			Resistance: mon.StatusResists.Slow,
		},
		{
			Status:     "threaten",
			Resistance: mon.StatusResists.Threaten,
		},
		{
			Status:     "zombie",
			Resistance: mon.StatusResists.Zombie,
		},
		{
			Status:     "power break",
			Resistance: mon.StatusResists.PowerBreak,
		},
		{
			Status:     "magic break",
			Resistance: mon.StatusResists.MagicBreak,
		},
		{
			Status:     "armor break",
			Resistance: mon.StatusResists.ArmorBreak,
		},
		{
			Status:     "mental break",
			Resistance: mon.StatusResists.MentalBreak,
		},
	})
}

func (mon *MonsterOld) getPosResists() []StatusResist {
	return filterStatusResists([]StatusResist{
		{
			Status:     "haste",
			Resistance: mon.StatusResists.Haste,
		},
		{
			Status:     "nul spells",
			Resistance: mon.StatusResists.NulSpells,
		},
		{
			Status:     "protect",
			Resistance: mon.StatusResists.Protect,
		},
		{
			Status:     "reflect",
			Resistance: mon.StatusResists.Reflect,
		},
		{
			Status:     "regen",
			Resistance: mon.StatusResists.Regen,
		},
		{
			Status:     "shell",
			Resistance: mon.StatusResists.Shell,
		},
	})
}

func (mon *MonsterOld) getMiscResists() []StatusResist {
	return filterStatusResists([]StatusResist{
		{
			Status:     "boost",
			Resistance: 0,
		},
		{
			Status:     "bribe",
			Resistance: 0,
		},
		{
			Status:     "defend",
			Resistance: 0,
		},
		{
			Status:     "distillers",
			Resistance: mon.StatusResists.Distiller,
		},
		{
			Status:     "guard",
			Resistance: 0,
		},
		{
			Status:     "life",
			Resistance: mon.StatusResists.Life,
		},
		{
			Status:     "scan",
			Resistance: mon.StatusResists.Scan,
		},
		{
			Status:     "sensor",
			Resistance: mon.StatusResists.Sensor,
		},
		{
			Status:     "sentinel",
			Resistance: 0,
		},
		{
			Status:     "shield",
			Resistance: 0,
		},
		{
			Status:     "curse",
			Resistance: 0,
		},
		{
			Status:     "percentage damage",
			Resistance: mon.getDemiResistance(),
		},
		{
			Status:     "slice",
			Resistance: 0,
		},
	})
}


func filterStatusResists(statusResists []StatusResist) []StatusResist {
	statusResistsFiltered := []StatusResist{}

	for _, status := range statusResists {
		if status.Resistance > 0 {
			statusResistsFiltered = append(statusResistsFiltered, status)
		}
	}

	return statusResistsFiltered
}

func (mon *MonsterOld) getDemiResistance() int {
	if val, ok := mon.ElemResists["gravity"].(float64); ok {
		switch val {
		case 0:
			return 254
		case 1:
			return 0
		}
	}

	return 0
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
