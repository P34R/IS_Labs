package model

type Period struct {
	Date string `json:"period"`
}

type BrandIndex struct {
	Id    int    `json:"-"`
	Month int    `json:"month,omitempty"`
	Year  int    `json:"year,omitempty"`
	Brand string `json:"brand"`
	Quant int    `json:"quantity"`
	Total int    `json:"total"`
}
