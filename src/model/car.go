package model

type Car struct {
	Id           int    `json:"id"`
	Brand        string `json:"brand"`
	Manufacturer string `json:"manufacturer"`
	Power        int    `json:"power"`
	Capacity     int    `json:"capacity"`
	Type         string `json:"type"`
	Price        int    `json:"price"`
}
