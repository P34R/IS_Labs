package store

import (
	"IS_Lab/src/model"
	"strconv"
)

type CarIndexRepository struct {
	Store *Store
}

func (r *CarIndexRepository) Create(month, year int) error {
	brands, err := r.Store.db.Query("SELECT \"brand\",count(*),sum(price) FROM \"carmoves\" where to_char(\"date\",'mm')=text($1) AND to_char(\"date\",'yyyy')=text($2) AND \"movetype\"=$3 group by \"brand\"", month, year, "-")
	if err != nil {
		return err
	}
	types, err := r.Store.db.Query("SELECT \"type\",count(*),sum(price) FROM \"carmoves\" where to_char(\"date\",'mm')=text($1) AND to_char(\"date\",'yyyy')=text($2) AND \"movetype\"=$3 group by \"type\"", month, year, "-")
	if err != nil {
		return err
	}
	for brands.Next() {
		var brand string
		var count, price int
		err = brands.Scan(&brand, &count, &price)
		if err != nil {
			return err
		}
		r.Store.db.QueryRow("INSERT INTO \"brandsales\" (\"periodmonth\", \"periodyear\",\"brand\",\"quantity\",\"totalprice\") VALUES ($1,$2,$3,$4,$5)", month, year, brand, count, price)
	}
	for types.Next() {
		var Type model.CarType
		var count, price int
		err = types.Scan(&Type, &count, &price)
		if err != nil {
			return err
		}
		r.Store.db.QueryRow("INSERT INTO \"typesales\" (\"periodmonth\", \"periodyear\",\"type\",\"quantity\",\"totalprice\") VALUES ($1,$2,$3,$4,$5)", month, year, Type, count, price)
	}

	return nil
}

func (r *CarIndexRepository) ReadType(month, year int) (*[]model.TypeIndex, error) {
	var res []model.TypeIndex

	rows, err := r.Store.db.Query("SELECT * FROM \"typesales\" WHERE \"periodmonth\"=$1 and \"periodyear\"=$2", month, year)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var m model.TypeIndex
		err = rows.Scan(&m.Id, &m.Month, &m.Year, &m.Type, &m.Quant, &m.Total)
		if err != nil {
			return nil, err
		}
		res = append(res, m)
	}
	return &res, nil
}
func (r *CarIndexRepository) ReadBrand(month, year int) (*[]model.BrandIndex, error) {
	var res []model.BrandIndex

	rows, err := r.Store.db.Query("SELECT * FROM \"brandsales\" WHERE \"periodmonth\"=$1 and \"periodyear\"=$2", month, year)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var m model.BrandIndex
		err = rows.Scan(&m.Id, &m.Brand, &m.Quant, &m.Total)
		if err != nil {
			return nil, err
		}
		res = append(res, m)
	}
	return &res, nil
}
func (r *CarIndexRepository) ReadBrandPeriod(dates []int) (*[]model.BrandIndex, error) {
	var res []model.BrandIndex
	//dates should contain 6 ints in this order => [dd,mm,yyyy,dd,mm,yyyy], where first 3 is start date, last 3 is end date
	startDate := strconv.Itoa(dates[0]) + "." + strconv.Itoa(dates[1]) + "." + strconv.Itoa(dates[2])
	endDate := strconv.Itoa(dates[3]) + "." + strconv.Itoa(dates[4]) + "." + strconv.Itoa(dates[5])
	rows, err := r.Store.db.Query("select \"brand\",count(*),sum(price) from carmoves where \"date\" between to_date($1,'dd.mm.yyyy') and to_date($2,'dd.mm.yyyy') group by \"brand\"", startDate, endDate)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var m model.BrandIndex
		err = rows.Scan(&m.Brand, &m.Quant, &m.Total)
		if err != nil {
			return nil, err
		}
		res = append(res, m)
	}
	return &res, nil
}

func (r *CarIndexRepository) ReadTypePeriod(dates []int) (*[]model.TypeIndex, error) {
	var res []model.TypeIndex
	//dates should contain 6 ints in this order => [dd,mm,yyyy,dd,mm,yyyy], where first 3 is start date, last 3 is end date
	startDate := strconv.Itoa(dates[0]) + "." + strconv.Itoa(dates[1]) + "." + strconv.Itoa(dates[2])
	endDate := strconv.Itoa(dates[3]) + "." + strconv.Itoa(dates[4]) + "." + strconv.Itoa(dates[5])
	rows, err := r.Store.db.Query("select \"type\",count(*),sum(price) from carmoves where \"date\" between to_date($1,'dd.mm.yyyy') and to_date($2,'dd.mm.yyyy') group by \"type\"", startDate, endDate)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var m model.TypeIndex
		err = rows.Scan(&m.Type, &m.Quant, &m.Total)
		if err != nil {
			return nil, err
		}
		m.Month = 0
		m.Year = 0
		res = append(res, m)
	}
	return &res, nil
}
