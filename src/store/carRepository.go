package store

import (
	"IS_Lab/src/model"
	"IS_Lab/src/parsers"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
)

type CarRepository struct {
	Store *Store
}

func (r *CarRepository) Create(m *model.Car) error {
	err := r.Store.db.QueryRow("INSERT INTO \"cars\" (\"brand\", \"model\",\"power\",\"capacity\",\"fuel\",\"type\",\"price\") VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING \"id\"", m.Brand, m.Model, m.Power, m.Capacity, m.Fuel, m.Type, m.Price).Scan(&m.Id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
func (r *CarRepository) Search(m *model.CarChars) ([]*model.Car, error) {
	first := true
	Request := "SELECT \"id\" FROM \"cars\" WHERE "
	if m.Brand != "" {
		first = false
		Request = Request + "\"brand\" = '" + m.Brand + "'"
	}
	if m.Model != "" {
		if !first {
			Request += " AND "
		}
		first = false
		Request = Request + "\"model\" = '" + m.Model + "'"
	}
	if m.Power != "" {
		if !first {
			Request += " AND "
		}
		code, ints, err := parsers.ParseIntegerSearch(m.Power)
		switch code {
		case parsers.ERROR:
			return nil, err
		case parsers.MORE:
			Request = Request + "\"power\" > " + strconv.Itoa(ints[0])
		case parsers.LESS:
			Request = Request + "\"power\" < " + strconv.Itoa(ints[0])
		case parsers.BETWEEN:
			Request = Request + "\"power\" BETWEEN " + strconv.Itoa(ints[0]) + " AND " + strconv.Itoa(ints[1])
		case parsers.EQUAL:
			Request = Request + "\"power\" = " + strconv.Itoa(ints[0])
		}
		first = false
	}
	if m.Capacity != "" {
		if !first {
			Request += " AND "
		}
		code, ints, err := parsers.ParseIntegerSearch(m.Capacity)
		switch code {
		case parsers.ERROR:
			return nil, err
		case parsers.MORE:
			Request = Request + "\"capacity\" > " + strconv.Itoa(ints[0])
		case parsers.LESS:
			Request = Request + "\"capacity\" < " + strconv.Itoa(ints[0])
		case parsers.BETWEEN:
			Request = Request + "\"capacity\" BETWEEN " + strconv.Itoa(ints[0]) + " AND " + strconv.Itoa(ints[1])
		case parsers.EQUAL:
			Request = Request + "\"capacity\" = " + strconv.Itoa(ints[0])
		}
		first = false
	}
	if m.Fuel != "" {
		if !first {
			Request += " AND "
		}
		Request = Request + "\"fuel\" = '" + m.Fuel + "'"
		first = false
	}
	if m.Type != "" {
		if !first {
			Request += " AND "
		}
		Request = Request + "\"type\" = '" + m.Type + "'"
		first = false
	}
	if m.Price != "" {
		if !first {
			Request += " AND "
		}
		code, ints, err := parsers.ParseIntegerSearch(m.Price)
		switch code {
		case parsers.ERROR:
			return nil, err
		case parsers.MORE:
			Request = Request + "\"price\" > " + strconv.Itoa(ints[0])
		case parsers.LESS:
			Request = Request + "\"price\" < " + strconv.Itoa(ints[0])
		case parsers.BETWEEN:
			Request = Request + "\"price\" BETWEEN " + strconv.Itoa(ints[0]) + " AND " + strconv.Itoa(ints[1])
		case parsers.EQUAL:
			Request = Request + "\"price\" = " + strconv.Itoa(ints[0])
		}
		first = false
	}
	if first {
		return nil, errors.New("error: empty filter")
	}
	var ret []*model.Car
	rows, err := r.Store.db.Query(Request)
	fmt.Println(rows)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var id int
		_ = rows.Scan(&id)
		if s, err := r.Read(id); err != sql.ErrNoRows {
			ret = append(ret, s)
		}
		if len(ret) == 10 {
			return ret, nil
		}
	}
	return ret, nil
}
func (r *CarRepository) Read(id int) (*model.Car, error) {
	var m model.Car
	err := r.Store.db.QueryRow("SELECT * FROM \"cars\" WHERE \"id\"=$1", id).Scan(&m.Id, &m.Brand, &m.Model, &m.Power, &m.Capacity, &m.Fuel, &m.Type, &m.Price, &m.IsSold)
	if err != nil {
		return nil, err
	}
	return &m, nil
}
func (r *CarRepository) Update(u *model.Car) error {
	n, err := r.Read(u.Id)
	if err != nil {
		return err
	}
	if u.Brand == "" {
		u.Brand = n.Brand
	}
	if u.Model == "" {
		u.Model = n.Model
	}
	if u.Power == 0 {
		u.Power = n.Power
	}
	if u.Capacity == 0 {
		u.Capacity = n.Capacity
	}
	if u.Fuel == "" {
		u.Fuel = n.Fuel
	}
	if u.Type == "" {
		u.Type = n.Type
	}
	if u.Price == 0 {
		u.Price = n.Price
	}
	_, err = r.Store.db.Exec("UPDATE \"cars\" SET \"brand\"=$1,\"model\"=$2,\"power\"=$3,\"capacity\"=$4,\"fuel\"=$5,\"type\"=$6,\"price\"=$7 WHERE \"id\"=$8", u.Brand, u.Model, u.Power, u.Capacity, u.Fuel, u.Type, u.Price, u.Id)
	if err != nil {
		return err
	}
	return nil
}
