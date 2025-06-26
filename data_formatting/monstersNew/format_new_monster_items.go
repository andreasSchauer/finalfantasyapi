package dataFormattingMonstersNew

import(
	old "github.com/andreasSchauer/finalfantasyapi/data_formatting/monstersOld"
)

type ItemsOld struct {
	ItemDropRate	float64 			`json:"item_drop_rate"`
	StealCommon		[]old.Item		`json:"steal_common"`
	StealRare		[]old.Item		`json:"steal_rare"`
	DropCommon		[]old.Item		`json:"drop_common"`
	DropRare		[]old.Item		`json:"drop_rare"`
	Bribe			[]old.Item		`json:"bribe"`
}


type ItemsNew struct {
	ItemDropRate			float64 			`json:"item_drop_rate"`
	StealCommon				*old.Item		`json:"steal_common"`
	StealRare				*old.Item		`json:"steal_rare"`
	DropCommon				*old.Item		`json:"drop_common"`
	DropRare				*old.Item		`json:"drop_rare"`
	SecondaryDropCommon		*old.Item		`json:"secondary_drop_common"`
	SecondaryDropRare		*old.Item		`json:"secondary_drop_rare"`
	Bribe					*old.Item		`json:"bribe"`
}



func (mon *MonsterOld) formatItems() *ItemsNew {
	if mon.Items == nil {
		return nil
	}

	items := ItemsNew{
		ItemDropRate: mon.Items.ItemDropRate,
		StealCommon: getItemData(mon.Items.StealCommon, false),
		StealRare: getItemData(mon.Items.StealRare, false),
		DropCommon: getItemData(mon.Items.DropCommon, false),
		DropRare: getItemData(mon.Items.DropRare, false),
		SecondaryDropCommon: getItemData(mon.Items.DropCommon, true),
		SecondaryDropRare: getItemData(mon.Items.DropRare, true),
		Bribe: getItemData(mon.Items.Bribe, false),
	}

	return &items
}


func getItemData(oldItem []old.Item, getSecondItem bool) *old.Item {
	if !getSecondItem {
		if len(oldItem) == 0 {
			return nil
		}
	
		if len(oldItem) >= 1 {
			return &oldItem[0]
		}
	}

	if getSecondItem {
		if len(oldItem) < 2 {
			return nil
		}
	
		if len(oldItem) == 2 {
			return &oldItem[1]
		}
	}

	return nil
}