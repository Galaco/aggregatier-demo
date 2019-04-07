package common

type God struct {
	Ability1 Ability `json:"Ability_1"`
	Ability2 Ability `json:"Ability_2"`
	Ability3 Ability `json:"Ability_3"`
	Ability4 Ability `json:"Ability_4"`
	Ability5 Ability `json:"Ability_5"`

	AttackSpeed                float32 `json:"AttackSpeed"`
	AttackSpeedPerLevel        float32 `json:"AttackSpeedPerLevel"`
	Cons                       string  `json:"Cons"`
	HP5PerLevel                float32 `json:"HP5PerLevel"`
	Health                     int     `json:"Health"`
	HealthPer5                 int     `json:"HealthPerFive"`
	HealthPerLevel             int     `json:"HealthPerLevel"`
	Lore                       string  `json:"Lore"`
	MP5PerLevel                float32 `json:"MP5PerLevel"`
	MagicProtection            float32 `json:"MagicProtection"`
	MagicProtectionPerLevel    float32 `json:"MagicProtectionPerLevel"`
	MagicalPower               float32 `json:"MagicalPower"`
	MagicalPowerPerLevel       float32 `json:"MagicalPowerPerLevel"`
	Mana                       int     `json:"Mana"`
	ManaPer5                   float32 `json:"ManaPerFive"`
	ManaPerLevel               float32 `json:"ManaPerLevel"`
	Name                       string  `json:"Name"`
	OnFreeRotation             string  `json:"OnFreeRotation"`
	Pantheon                   string  `json:"Pantheon"`
	PhysicalPower              int     `json:"PhysicalPower"`
	PhysicalPowerPerLevel      float32 `json:"PhysicalPowerPerLevel"`
	PhysicalProtection         float32 `json:"PhysicalProtection"`
	PhysicalProtectionPerLevel float32 `json:"PhysicalProtectionPerLevel"`
	Pros                       string  `json:"Pros"`
	Roles                      string  `json:"Roles"`
	Speed                      int     `json:"Speed"`
	Title                      string  `json:"Title"`
	Type                       string  `json:"Type"`
	CardUrl                    string  `json:"godCard_URL"`
	IconUrl                    string  `json:"godIcon_URL"`
	Id                         int     `json:"id"`
	LatestGod                  string  `json:"latestGod"`
}

type Ability struct {
	Description struct {
		ItemDescription ItemDescription `json:"itemDescription"`
	} `json:"Description"`
	Id      int    `json:"Id"`
	Summary string `json:"Summary"`
	Url     string `json:"URL"`
}

type ItemDescription struct {
	Cooldown             string     `json:"cooldown"`
	Cost                 string     `json:"cost"`
	Description          string     `json:"description"`
	MenuItems            []KeyValue `json:"menuitems"`
	RankItems            []KeyValue `json:"rankitems"`
	SecondaryDescription string     `json:"secondaryDescription"`
}

type KeyValue struct {
	Description string `json:"description"`
	Value       string `json:"value"`
}
