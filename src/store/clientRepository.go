package store

import (
	"IS_Lab/src/model"
)

type ClientRepository struct {
	Store *Store
}

func (r *ClientRepository) Create(m *model.Client) error {
	err := r.Store.db.QueryRow("INSERT INTO \"clients\" (\"first_name\", \"last_name\",\"phone\") VALUES ($1,$2,$3) RETURNING \"id\"", m.FirstName, m.LastName, m.Phone).Scan(&m.Id)
	if err != nil {
		return err
	}
	return nil
}

func (r *ClientRepository) Read(id int) (*model.Client, error) {
	var m model.Client
	err := r.Store.db.QueryRow("SELECT * FROM \"clients\" WHERE \"id\"=$1", id).Scan(&m.Id, &m.FirstName, &m.LastName, &m.Phone)
	if err != nil {
		return nil, err
	}
	return &m, nil
}
