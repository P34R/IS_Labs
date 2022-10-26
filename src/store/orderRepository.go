package store

import (
	"IS_Lab/src/model"
)

type OrderRepository struct {
	Store *Store
}

func (r *OrderRepository) Create(m *model.Order) error {
	err := r.Store.db.QueryRow("INSERT INTO \"orders\" (\"date\", \"client_id\",\"car_id\", \"status\") VALUES ($1,$2,$3,$4) RETURNING \"id\"", m.Date, m.ClientId, m.CarId, m.Status).Scan(&m.Id)
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) Read(id int) (*model.Order, error) {
	var m model.Order
	err := r.Store.db.QueryRow("SELECT * FROM \"orders\" WHERE \"id\"=$1", id).Scan(&m.Id, &m.Date, &m.ClientId, &m.CarId, &m.Status)
	if err != nil {
		return nil, err
	}
	return &m, nil
}
