package model

type Status string

const (
	DONE     Status = "done"
	PROGRESS        = "in progress"
	CANCELED        = "canceled"
	EMPTY           = ""
)

type Order struct {
	Id       int    `json:"-"`
	Date     string `json:"date,omitempty"`
	ClientId int    `json:"client_id"`
	CarId    int    `json:"car_id"`
	Status   Status `json:"status,omitempty"`
}
