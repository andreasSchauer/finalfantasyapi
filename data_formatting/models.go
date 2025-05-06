package dataFormatting

type WpnAbilityOld struct {
	Ability    string `json:"ability"`
	Characters string `json:"characters,omitempty"`
}
type ArmorAbilityOld struct {
	Ability string `json:"ability"`
}

type EquipmentOld struct {
	DropRate          float64            `json:"drop_rate"`
	SlotsAmount       string             `json:"slots_amount"`
	AttachedAbilities string             `json:"attached_abilities"`
	WpnAbilities      []WpnAbilityOld    `json:"wpn_abilities"`
	ArmorAbilities    []ArmorAbilityOld  `json:"armor_abilities"`
}

type StatsOld struct {
	HP         []any `json:"hp"`
	MP         int   `json:"mp"`
	Strength   int   `json:"strength"`
	Defence    int   `json:"defence"`
	Magic      int   `json:"magic"`
	MagDefence int   `json:"mag_defence"`
	Agility    int   `json:"agility"`
	Luck       int   `json:"luck"`
	Evasion    int   `json:"evasion"`
	Accuracy   int   `json:"accuracy"`
}

type StatusResistsOld struct {
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
	ArmorBreak   int   `json:"armor_break"`
	MentalBreak  int   `json:"mental_break"`
	Threaten     int   `json:"threaten"`
	Death        int   `json:"death"`
	Provoke      int   `json:"provoke"`
	Doom         []int `json:"doom"`
	NulMagic     int   `json:"nul_magic"`
	Shell        int   `json:"shell"`
	Protect      int   `json:"protect"`
	Reflect      int   `json:"reflect"`
	Haste        int   `json:"haste"`
	Regen        int   `json:"regen"`
	Distiller    int   `json:"distiller"`
	Sensor       int   `json:"sensor"`
	Scan         int   `json:"scan"`
	Delay        int   `json:"delay"`
	Eject        int   `json:"eject"`
	Berserk      int   `json:"berserk"`
	ZanmatoLevel int   `json:"zanmato_level"`
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

type Item struct {
	Item   string `json:"item"`
	Amount int    `json:"amount"`
}

type Characters struct {
	KimahriAndAuron   bool `json:"kimahri_and_auron"`
	YunaAndLulu       bool `json:"yuna_and_lulu"`
	ExceptYunaAndLulu bool `json:"except_yuna_and_lulu"`
}
type WpnAbility struct {
	Ability    string     `json:"ability"`
	Characters Characters `json:"characters"`
}
type ArmorAbility struct {
	Ability string `json:"ability"`
}
type Equipment struct {
	DropRate             float64        `json:"drop_rate"`
	MinSlotsAmount       int            `json:"min_slots_amount"`
	MaxSlotsAmount       int            `json:"max_slots_amount"`
	MinAttachedAbilities int            `json:"min_attached_abilities"`
	MaxAttachedAbilities int            `json:"max_attached_abilities"`
	WpnAbilities         []WpnAbility   `json:"wpn_abilities"`
	ArmorAbilities       []ArmorAbility `json:"armor_abilities"`
}
type Stats struct {
	HP             int  `json:"hp"`
	OverkillDamage *int `json:"overkill_damage"`
	MP             int  `json:"mp"`
	Strength       int  `json:"strength"`
	Defence        int  `json:"defence"`
	Magic          int  `json:"magic"`
	MagDefence     int  `json:"mag_defence"`
	Agility        int  `json:"agility"`
	Luck           int  `json:"luck"`
	Evasion        int  `json:"evasion"`
	Accuracy       int  `json:"accuracy"`
}

type ElemResist struct {
	Element  string `json:"element"`
	Affinity string `json:"affinity"`
}

type StatusResist struct {
	Status     string `json:"status"`
	Resistance int    `json:"resistance"`
}

type StatusResists struct {
	Resistances    []StatusResist `json:"resistances"`
	PoisonRate     *float64       `json:"poison_rate"`
	DoomCountdown  *int           `json:"doom_countdown"`
	ZanmatoLevel   int            `json:"zanmato_level"`
	IsImmuneToLife bool           `json:"is_immune_to_life"`
}

type Monster struct {
	Location        []string          `json:"location"`
	Species         string            `json:"species"`
	IsReoccurring   bool              `json:"is_reoccurring"`
	IsCatchable     bool              `json:"is_catchable"`
	IsBoss          bool              `json:"is_boss"`
	IsZombie        bool              `json:"is_zombie"`
	IsTough         bool              `json:"is_tough"`
	IsHeavy         bool              `json:"is_heavy"`
	IsArmored       bool              `json:"is_armored"`
	HasOverdrive    bool              `json:"has_overdrive"`
	ImmuneToPhysDmg bool              `json:"immune_to_phys_dmg"`
	ImmuneToMagDmg  bool              `json:"immune_to_mag_dmg"`
	Allies          []string          `json:"allies"`
	Ap              int               `json:"ap"`
	ApOverkill      int               `json:"ap_overkill"`
	Gil             int               `json:"gil"`
	RonsoRage       []string          `json:"ronso_rage"`
	Items           map[string][]Item `json:"items"`
	Equipment       Equipment         `json:"equipment"`
	Stats           Stats             `json:"stats"`
	ElemResists     []ElemResist      `json:"elem_resists"`
	StatusResists   StatusResists     `json:"status_resists"`
}
