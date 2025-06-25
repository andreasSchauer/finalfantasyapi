package dataFormattingMonsters

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
)

type Monster struct {
	Name           	string          `json:"monster_name"`
	Species        	string          `json:"species"`
	IsReoccurring  	bool            `json:"is_reoccurring"`
	IsCatchable    	bool            `json:"can_be_caught"`
	IsBoss         	bool            `json:"is_boss"`
	IsUnderwater   	bool            `json:"is_underwater"`
	Properties     	[]string        `json:"properties"`
	HasOverdrive   	bool            `json:"has_overdrive"`
	Ap             	int             `json:"ap"`
	ApOverkill     	int             `json:"ap_overkill"`
	OverkillDamage 	*int            `json:"overkill_damage"`
	Gil            	int             `json:"gil"`
	RonsoRage      	[]string        `json:"ronso_rage"`
	PoisonRate      *float64       	`json:"poison_rate"`
	DoomCountdown   *int           	`json:"doom_countdown"`
	ThreatenCounter int            	`json:"threaten_counter"`
	Items          	Items           `json:"items"`
	Equipment      	Equipment       `json:"equipment"`
	Stats          	[]Stat          `json:"stats"`
	ElemResists    	[]ElemResist    `json:"elem_resists"`
	StatusResists  	StatusResists 	`json:"status_resists"`
}

type MonsterOld struct {
	Location      []string         `json:"location"`
	IsReoccurring bool             `json:"is_reoccurring"`
	IsCatchable   bool             `json:"is_catchable"`
	IsBoss        bool             `json:"is_boss"`
	IsZombie      bool             `json:"is_zombie"`
	Allies        []string         `json:"allies"`
	Ap            []int            `json:"ap"`
	Gil           int              `json:"gil"`
	RonsoRage     []string         `json:"ronso_rage"`
	Items         map[string][]any `json:"items"`
	Equipment     EquipmentOld     `json:"equipment"`
	Stats         StatsOld         `json:"stats"`
	ElemResists   map[string]any   `json:"elem_resists"`
	StatusResists StatusResistsOld `json:"stat_resists"`
}

func FormatMonsterJson() error {
	const srcPath = "./data_old/monsters.json"
	const destPath = "./data/monsters_copy.json"

	file, err := os.Open(srcPath)
	if err != nil {
		return fmt.Errorf("couldn't open file: %v", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("couldn't read file: %v", err)
	}

	monstersOld := make(map[string]MonsterOld)
	err = json.Unmarshal(data, &monstersOld)
	if err != nil {
		return fmt.Errorf("couldn't parse JSON: %v", err)
	}

	var monstersNew []Monster

	for name, monData := range monstersOld {
		stats, overkillDamage := monData.formatStats()
		statusResistsData := monData.formatStatusResists()

		mon := Monster{
			Name:           name,
			Species:        "",
			IsReoccurring:  monData.IsReoccurring,
			IsCatchable:    monData.IsCatchable,
			IsBoss:         monData.IsBoss,
			Properties:     []string{},
			Ap:             monData.Ap[0],
			ApOverkill:     monData.Ap[1],
			OverkillDamage: overkillDamage,
			Gil:            monData.Gil,
			RonsoRage:      monData.RonsoRage,
			PoisonRate: 	statusResistsData.PoisonRate,
			DoomCountdown: 	statusResistsData.DoomCountdown,
			ThreatenCounter: statusResistsData.ThreatenCounter,
			Items:          monData.formatItems(),
			Equipment:      monData.formatEquipment(),
			Stats:          stats,
			ElemResists:    monData.formatElemResist(),
			StatusResists:  statusResistsData.StatusResists,
		}

		monstersNew = append(monstersNew, mon)
	}

	sort.SliceStable(monstersNew, func(i, j int) bool { return monstersNew[i].Name < monstersNew[j].Name })

	monstersJson := MonstersData{
		MonsterData: monstersNew,
	}

	json, err := json.MarshalIndent(monstersJson, "", "    ")
	if err != nil {
		return fmt.Errorf("couldn't encode JSON: %v", err)
	}

	os.WriteFile(destPath, json, 0666)

	return nil
}

type MonstersData struct {
	MonsterData []Monster `json:"monsters_data"`
}
