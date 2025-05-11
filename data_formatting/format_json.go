package dataFormatting

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
)


func FormatJson() error {
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






