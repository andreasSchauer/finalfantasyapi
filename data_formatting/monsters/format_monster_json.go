package dataFormattingMonsters

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
)



type Monster struct {
	Name            string            `json:"monster_name"`
	Location        []string          `json:"location"`
	Species         string            `json:"species"`
	IsReoccurring   bool              `json:"is_reoccurring"`
	IsCatchable     bool              `json:"can_be_caught"`
	IsBoss          bool              `json:"is_boss"`
	IsZombie        bool              `json:"is_zombie"`
	IsTough         bool              `json:"is_tough"`
	IsHeavy         bool              `json:"is_heavy"`
	IsArmored       bool              `json:"is_armored"`
	IsUnderwater    bool              `json:"is_underwater"`
	HasOverdrive    bool              `json:"has_overdrive"`
	ImmuneToPhysDmg bool              `json:"immune_to_phys_dmg"`
	ImmuneToMagDmg  bool              `json:"immune_to_mag_dmg"`
	Allies          []string          `json:"allies"`
	Ap              int               `json:"ap"`
	ApOverkill      int               `json:"ap_overkill"`
	OverkillDamage  *int              `json:"overkill_damage"`
	Gil             int               `json:"gil"`
	RonsoRage       []string          `json:"ronso_rage"`
	Items           map[string][]Item `json:"items"`
	Equipment       Equipment         `json:"equipment"`
	Stats           Stats             `json:"stats"`
	ElemResists     []ElemResist      `json:"elem_resists"`
	StatusResists   StatusResists     `json:"status_resists"`
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
	const filepath = "./data_old/monsters.json"

	file, err := os.Open(filepath)
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
		mon := Monster{
			Name: name,
			Location: monData.Location,
			Species: "",
			IsReoccurring: monData.IsReoccurring,
			IsCatchable: monData.IsCatchable,
			IsBoss: monData.IsBoss,
			IsZombie: monData.IsZombie,
			Allies: monData.formatAllies(name),
			Ap: monData.Ap[0],
			ApOverkill: monData.Ap[1],
			OverkillDamage: overkillDamage,
			Gil: monData.Gil,
			RonsoRage: monData.RonsoRage,
			Items: monData.formatItems(),
			Equipment: monData.formatEquipment(),
			Stats: stats,
			ElemResists: monData.formatElemResist(),
			StatusResists: monData.formatStatusResists(),
		}

		monstersNew = append(monstersNew, mon)
	}

	sort.SliceStable(monstersNew, func(i, j int) bool { return monstersNew[i].Name < monstersNew[j].Name })

	monstersJson := struct{
		MonsterData	[]Monster	`json:"monsters_data"`
	}{
		MonsterData: monstersNew,
	}

	json, err := json.MarshalIndent(monstersJson, "", "    ")
	if err != nil {
		return fmt.Errorf("couldn't encode JSON: %v", err)
	}

	fmt.Println(string(json))

	return nil
}






