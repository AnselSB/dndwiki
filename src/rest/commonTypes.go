package rest

type Damage struct {
	DamageType
	DamageAtCharacterLevel
}

type DamageType struct {
	Name string `json:"name"`
}

type DamageAtCharacterLevel struct {
	First       string `json:"1"`
	Fifth       string `json:"5"`
	Eleventh    string `json:"11"`
	Seventeenth string `json:"17"`
}

type EquipmentCategory struct {
	Index string `json:"index"`
	Name  string `json:"name"`
	url   string `json"url"`
}

type Rarity struct {
	Name string `json:"name"`
}

type Variant struct {
	Name string `json:"name"`
}
