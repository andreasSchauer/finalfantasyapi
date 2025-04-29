package data_formatting




type WpnAbilitiesOld struct {
	Ability    string `json:"ability"`
	Characters string `json:"characters,omitempty"`
}
type ArmourAbilitiesOld struct {
	Ability string `json:"ability"`
}
type EquipmentOld struct {
	DropRate          float64           `json:"drop_rate"`
	SlotsAmount       string            `json:"slots_amount"`
	AttachedAbilities string            `json:"attached_abilities"`
	WpnAbilities      []WpnAbilitiesOld    `json:"wpn_abilities"`
	ArmourAbilities   []ArmourAbilitiesOld `json:"armour_abilities"`
}
type StatsOld struct {
	Hp         []int `json:"hp"`
	Mp         int   `json:"mp"`
	Strength   int   `json:"strength"`
	Defence    int   `json:"defence"`
	Magic      int   `json:"magic"`
	MagDefence int   `json:"mag_defence"`
	Agility    int   `json:"agility"`
	Luck       int   `json:"luck"`
	Evasion    int   `json:"evasion"`
	Accuracy   int   `json:"accuracy"`
}
type ElemResistsOld struct {
	Fire      float64     `json:"fire"`
	Lightning float64     `json:"lightning"`
	Water     float64     `json:"water"`
	Ice       float64     `json:"ice"`
	Holy      float64 `json:"holy"`
	Gravity   int     `json:"gravity"`
}
type StatResistsOld struct {
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
	ArmourBreak  int   `json:"armour_break"`
	MentalBreak  int   `json:"mental_break"`
	Threaten     int   `json:"threaten"`
	Death        int   `json:"death"`
	Provoke      int   `json:"provoke"`
	Doom         []int `json:"doom"`
	Delay        int   `json:"delay"`
	Eject        int   `json:"eject"`
	ZanmatoLevel int   `json:"zanmato_level"`
}
type MonsterOld struct {
	Location      []string    `json:"location"`
	IsReoccurring bool        `json:"is_reoccurring"`
	IsCatchable   bool        `json:"is_catchable"`
	IsBoss        bool        `json:"is_boss"`
	IsZombie      bool        `json:"is_zombie"`
	Allies        []string    `json:"allies"`
	Ap            []int       `json:"ap"`
	Gil           int         `json:"gil"`
	RonsoRage     []string    `json:"ronso_rage"`
	Items         map[string][]any      `json:"items"`
	Equipment     EquipmentOld   `json:"equipment"`
	Stats         StatsOld       `json:"stats"`
	ElemResists   ElemResistsOld `json:"elem_resists"`
	StatResists   StatResistsOld `json:"stat_resists"`
}




type Item struct {
	Item   string `json:"item"`
	Amount int    `json:"amount"`
}

type Characters struct {
	KimahriAndAuron   bool `json:"kimahri_and_auron"`
	YunaAndLulu       bool `json:"yuna_and_lulu"`
	ExceptYunaAndLulu bool `json:"except_yuna_and_lulu"`
}
type WpnAbilities struct {
	Ability    string     `json:"ability"`
	Characters Characters `json:"characters"`
}
type ArmourAbilities struct {
	Ability string `json:"ability"`
}
type Equipment struct {
	DropRate             float64           `json:"drop_rate"`
	MinSlotsAmount       int               `json:"min_slots_amount"`
	MaxSlotsAmount       int               `json:"max_slots_amount"`
	MinAttachedAbilities int               `json:"min_attached_abilities"`
	MaxAttachedAbilities int               `json:"max_attached_abilities"`
	WpnAbilities         []WpnAbilities    `json:"wpn_abilities"`
	ArmourAbilities      []ArmourAbilities `json:"armour_abilities"`
}
type Stats struct {
	Hp             int `json:"hp"`
	OverkillDamage int `json:"overkill_damage"`
	Mp             int `json:"mp"`
	Strength       int `json:"strength"`
	Defence        int `json:"defence"`
	Magic          int `json:"magic"`
	MagDefence     int `json:"mag_defence"`
	Agility        int `json:"agility"`
	Luck           int `json:"luck"`
	Evasion        int `json:"evasion"`
	Accuracy       int `json:"accuracy"`
}
type ElemResists struct {
	Fire      string `json:"fire"`
	Lightning string `json:"lightning"`
	Water     string `json:"water"`
	Ice       string `json:"ice"`
	Holy      string `json:"holy"`
}
type StatResists struct {
	Silence          int     `json:"silence"`
	Sleep            int     `json:"sleep"`
	Dark             int     `json:"dark"`
	Poison           int     `json:"poison"`
	PoisonPercentage float64 `json:"poison_percentage"`
	Petrify          int     `json:"petrify"`
	Slow             int     `json:"slow"`
	Zombie           int     `json:"zombie"`
	Life             int     `json:"life"`
	PowerBreak       int     `json:"power_break"`
	MagicBreak       int     `json:"magic_break"`
	ArmourBreak      int     `json:"armour_break"`
	MentalBreak      int     `json:"mental_break"`
	Threaten         int     `json:"threaten"`
	Death            int     `json:"death"`
	Provoke          int     `json:"provoke"`
	Doom             int     `json:"doom"`
	DoomCountdown    int     `json:"doom_countdown"`
	Demi             int     `json:"demi"`
	Delay            int     `json:"delay"`
	Eject            int     `json:"eject"`
	ZanmatoLevel     int     `json:"zanmato_level"`
}
type Monster struct {
	Location      []string    		 `json:"location"`
	IsReoccurring bool        		 `json:"is_reoccurring"`
	IsCatchable   bool        		 `json:"is_catchable"`
	IsBoss        bool        		 `json:"is_boss"`
	IsZombie      bool        		 `json:"is_zombie"`
	Allies        []string    		 `json:"allies"`
	Ap            int         		 `json:"ap"`
	ApOverkill    int         		 `json:"ap_overkill"`
	Gil           int         		 `json:"gil"`
	RonsoRage     []string    		 `json:"ronso_rage"`
	Items         map[string][]Item  `json:"items"`
	Equipment     Equipment   		 `json:"equipment"`
	Stats         Stats       		 `json:"stats"`
	ElemResists   ElemResists 		 `json:"elem_resists"`
	StatResists   StatResists 		 `json:"stat_resists"`
}