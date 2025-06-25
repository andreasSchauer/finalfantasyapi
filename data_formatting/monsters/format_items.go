package dataFormattingMonsters



type Item struct {
	Item   string `json:"item"`
	Amount int    `json:"amount"`
}


type Items struct {
	StealCommon		[]Item		`json:"steal_common"`
	StealRare		[]Item		`json:"steal_rare"`
	DropCommon		[]Item		`json:"drop_common"`
	DropRare		[]Item		`json:"drop_rare"`
	Bribe			[]Item		`json:"bribe"`
	StealGil		int			`json:"steal_gil"`
}



func (mon *MonsterOld) formatItems() Items {
	return Items{
		StealCommon: mon.getItems("steal_common"),
		StealRare: mon.getItems("steal_rare"),
		DropCommon: mon.getItems("drop_common"),
		DropRare: mon.getItems("drop_rare"),
		Bribe: mon.getItems("bribe"),
	}
}



func (mon *MonsterOld) getItems(key string) []Item {
	itemData := mon.Items[key]

	if itemData == nil {
		return []Item{}
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