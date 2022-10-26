package store

import (
	"IS_Lab/src/model"
)

type CarRepository struct {
	Store *Store
}

func (r *CarRepository) Create(m *model.Car) error {
	err := r.Store.db.QueryRow("INSERT INTO \"cars\" (\"brand\", \"manufacturer\",\"power\",\"capacity\",\"type\",\"price\") VALUES ($1,$2,$3,$4,$5,$6) RETURNING \"id\"", m.Brand, m.Manufacturer, m.Power, m.Capacity, m.Type, m.Price).Scan(&m.Id)
	if err != nil {
		return err
	}
	return nil
}

func (r *CarRepository) Read(id int) (*model.Car, error) {
	var m model.Car
	err := r.Store.db.QueryRow("SELECT * FROM \"cars\" WHERE \"id\"=$1", id).Scan(&m.Id, &m.Brand, &m.Manufacturer, &m.Power, &m.Capacity, &m.Type, &m.Price)
	if err != nil {
		return nil, err
	}
	return &m, nil
}
