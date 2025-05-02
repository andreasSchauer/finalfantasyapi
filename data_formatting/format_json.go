package dataFormatting

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)


func FormatJson() error {
	const filepath = "./data_formatting/monster_old_formatted.json"

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

	monstersNew := make(map[string]Monster)
	
	for name, monData := range monstersOld {
		monstersNew[name] = Monster{
			Location: monData.Location,
			IsReoccurring: monData.IsReoccurring,
			IsCatchable: monData.IsCatchable,
			IsBoss: monData.IsBoss,
			IsZombie: monData.IsZombie,
			Allies: monData.formatAllies(name),
			Ap: monData.Ap[0],
			ApOverkill: monData.Ap[1],
			Gil: monData.Gil,
			RonsoRage: monData.RonsoRage,
			Items: monData.formatItems(),
			Equipment: monData.formatEquipment(),
			Stats: monData.formatStats(),
			ElemResists: monData.formatElemResist(),
			StatusResists: monData.formatStatusResists(),
		}
	}

	json, err := json.MarshalIndent(monstersNew, "", " ")
	if err != nil {
		return fmt.Errorf("couldn't encode JSON: %v", err)
	}

	fmt.Println(string(json))

	return nil
}






