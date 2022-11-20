package model

type CarType string
type FuelType string

const (
	A CarType = "A"
	B         = "B"
	C         = "C"
	D         = "D"
	E         = "E"
	F         = "F"
	S         = "S"
	M         = "M"
	J         = "J"
)
const (
	GASOLINE FuelType = "gasoline"
	DIESEL            = "diesel"
	ELECTRIC          = "electric"
	BIOFUEL           = "biofuel"
	GAS               = "gas"
)

type Car struct {
	Id       int      `json:"-"`
	Brand    string   `json:"brand"`
	Model    string   `json:"model"`
	Power    int      `json:"power"`
	Capacity int      `json:"capacity"`
	Fuel     FuelType `json:"fuel"`
	Type     CarType  `json:"type"`
	Price    int      `json:"price"`
	IsSold   bool     `json:"sold,omitempty"`
}
type CarChars struct {
	Brand    string `json:"brand,omitempty"`
	Model    string `json:"model,omitempty"`
	Power    string `json:"power,omitempty"`
	Capacity string `json:"capacity,omitempty"`
	Fuel     string `json:"fuel,omitempty"`
	Type     string `json:"type,omitempty"`
	Price    string `json:"price,omitempty"`
}
