package model

type Order struct {
	Id       int    `json:"id"`
	Date     string `json:"date"`
	ClientId int    `json:"client_id"`
	CarId    int    `json:"car_id"`
	Status   bool   `json:"status"`
}
