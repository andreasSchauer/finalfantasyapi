package dataFormattingMonstersNew

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	old "github.com/andreasSchauer/finalfantasyapi/data_formatting/monstersOld"
)


type MonstersDataOld struct {
	MonsterData []MonsterOld `json:"monsters_data"`
}

type MonstersDataNew struct {
	MonsterData []MonsterNew `json:"monsters_data"`
}


type MonsterOld struct {
	Name            string        		`json:"name"`
	Species         string        		`json:"species"`
	IsReoccurring   bool          		`json:"is_reoccurring"`
	IsCatchable     bool          		`json:"can_be_captured"`
	IsBoss          bool          		`json:"is_boss"`
	HasOverdrive    bool          		`json:"has_overdrive"`
	IsUnderwater    bool          		`json:"is_underwater"`
	Properties      []string      		`json:"properties"`
	Ap              int           		`json:"ap"`
	ApOverkill      int           		`json:"ap_overkill"`
	OverkillDamage  *int          		`json:"overkill_damage"`
	Gil             int           		`json:"gil"`
	StealGil        int           		`json:"steal_gil"`
	RonsoRage       []string      		`json:"ronso_rage"`
	PoisonRate      *float64      		`json:"poison_rate"`
	DoomCountdown   *int          		`json:"doom_countdown"`
	ThreatenCounter *int          		`json:"threaten_counter"`
	ZanmatoLevel    int           		`json:"zanmato_level"`
	Stats           []old.Stat        	`json:"stats"`
	Items           *ItemsOld        	`json:"items"`
	Equipment       *old.Equipment    	`json:"equipment"`
	ElemResists     []old.ElemResist  	`json:"elem_resists"`
	StatusResists   old.StatusResists 	`json:"status_resists"`
	AlteredStates   []AlteredStateOld   `json:"altered_states"`
}



type MonsterNew struct {
	Name            string        			`json:"name"`
	Species         string        			`json:"species"`
	IsReoccurring   bool          			`json:"is_reoccurring"`
	IsCatchable     bool          			`json:"can_be_captured"`
	IsBoss          bool          			`json:"is_boss"`
	HasOverdrive    bool          			`json:"has_overdrive"`
	IsUnderwater    bool          			`json:"is_underwater"`
	Properties      []string     			`json:"properties"`
	Ap              int          			`json:"ap"`
	ApOverkill      int          			`json:"ap_overkill"`
	OverkillDamage  *int          			`json:"overkill_damage"`
	Gil             int           			`json:"gil"`
	StealGil        int           			`json:"steal_gil"`
	RonsoRage       []string      			`json:"ronso_rage"`
	PoisonRate      *float64      			`json:"poison_rate"`
	DoomCountdown   *int          			`json:"doom_countdown"`
	ThreatenCounter *int          			`json:"threaten_counter"`
	ZanmatoLevel    int           			`json:"zanmato_level"`
	Stats           []old.Stat        		`json:"stats"`
	Items           *ItemsNew        		`json:"items"`
	Equipment       *old.Equipment    		`json:"equipment"`
	ElemResists     []old.ElemResist 		`json:"elem_resists"`
	StatusResists   StatusResistsNew 		`json:"status_resists"`
	AlteredStates   []AlteredStateNew      	`json:"altered_states"`
}



func FormatNewMonsterJson() error {
	const srcPath = "./data/monsters_new.json"
	const destPath = "./data/monsters_new_copy.json"

	file, err := os.Open(srcPath)
	if err != nil {
		return fmt.Errorf("couldn't open file: %v", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("couldn't read file: %v", err)
	}

	monsterData := MonstersDataOld{}
	err = json.Unmarshal(data, &monsterData)
	if err != nil {
		return fmt.Errorf("couldn't parse JSON: %v", err)
	}

	var monstersNew []MonsterNew

	for _, monData := range monsterData.MonsterData {
		mon := MonsterNew{
			Name:            monData.Name,
			Species:         monData.Species,
			IsReoccurring:   monData.IsReoccurring,
			IsCatchable:     monData.IsCatchable,
			IsBoss:          monData.IsBoss,
			HasOverdrive: 	 monData.HasOverdrive,
			IsUnderwater:    monData.IsUnderwater,
			Properties:      monData.Properties,
			Ap:              monData.Ap,
			ApOverkill:      monData.ApOverkill,
			OverkillDamage:  monData.OverkillDamage,
			Gil:             monData.Gil,
			RonsoRage:       monData.RonsoRage,
			PoisonRate:      monData.PoisonRate,
			DoomCountdown:   monData.DoomCountdown,
			ThreatenCounter: monData.ThreatenCounter,
			ZanmatoLevel:    monData.ZanmatoLevel,
			Items:           monData.formatItems(),
			Equipment:       monData.Equipment,
			Stats:           monData.Stats,
			ElemResists:     monData.ElemResists,
			StatusResists:   *formatStatusResists(monData.StatusResists),
			AlteredStates:   monData.getAlteredStates(),
		}

		monstersNew = append(monstersNew, mon)
	}

	monstersJson := MonstersDataNew{
		MonsterData: monstersNew,
	}

	json, err := json.MarshalIndent(monstersJson, "", "    ")
	if err != nil {
		return fmt.Errorf("couldn't encode JSON: %v", err)
	}

	os.WriteFile(destPath, json, 0666)

	return nil
}