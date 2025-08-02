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
