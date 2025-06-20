package dataFormatting

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