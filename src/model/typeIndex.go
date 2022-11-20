package model

type TypeIndex struct {
	Id    int     `json:"-"`
	Month int     `json:"month,omitempty"`
	Year  int     `json:"year,omitempty"`
	Type  CarType `json:"type"`
	Quant int     `json:"quantity"`
	Total int     `json:"total"`
}
