package model

type Client struct {
	Id        int    `json:"-"`
	FirstName string `json:"name"`
	LastName  string `json:"surname"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
}
