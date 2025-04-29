package data_formatting

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
			Allies: monData.Allies,
			Ap: monData.Ap[0],
			ApOverkill: monData.Ap[1],
			Gil: monData.Gil,
			RonsoRage: monData.RonsoRage,
			Items: monData.formatItems(),
		}
	}

	json, err := json.MarshalIndent(monstersNew, "", " ")
	if err != nil {
		return fmt.Errorf("couldn't encode JSON: %v", err)
	}

	fmt.Println(string(json))

	return nil
}


func (mon *MonsterOld) formatItems() map[string][]Item {
	itemMap := make(map[string][]Item)
	keys := []string{"steal_common", "steal_rare", "drop_common", "drop_rare", "bribe"}

	for _, key := range keys {
		itemMap[key] = mon.getItems(key)
	}

	return itemMap
}


func (mon *MonsterOld) getItems(key string) []Item {
	itemData := mon.Items[key]

	if itemData == nil {
		return nil
	}

	var items []Item

	switch itemData[0].(type) {
	case string:
		items = append(items, formatItem(itemData))
	case []any:
		for _, data := range(itemData) {
			d, _ := data.([]any)
			items = append(items, formatItem(d))
		}
	}

	return items
}


func formatItem(itemData []any) Item {
	itemName, _ := itemData[0].(string)
	itemAmountFloat, _ := itemData[1].(float64)
	itemAmount := int(itemAmountFloat)

	return Item{
		Item: itemName,
		Amount: itemAmount,
	}
}