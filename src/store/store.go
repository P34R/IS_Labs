package store

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Store struct {
	url    string
	db     *sql.DB
	car    *CarRepository
	client *ClientRepository
	order  *OrderRepository
}

func NewStore() *Store {

	store := &Store{
		url: "host=localhost dbname=InfoSys port=5432 user=postgres password=admin sslmode=disable",
	}

	return store
}

func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.url)
	if err != nil {
		fmt.Print("Open error ")
		return err
	}
	if err := db.Ping(); err != nil {
		fmt.Print("Ping error ")
		return err
	}
	s.db = db
	return nil
}

func (s *Store) Close() {
	s.db.Close()
}

func (s *Store) Car() *CarRepository {
	if s.car == nil {
		s.car = &CarRepository{
			Store: s,
		}
	}
	return s.car
}
func (s *Store) Client() *ClientRepository {
	if s.client == nil {
		s.client = &ClientRepository{
			Store: s,
		}
	}
	return s.client
}
func (s *Store) Order() *OrderRepository {
	if s.order == nil {
		s.order = &OrderRepository{
			Store: s,
		}
	}
	return s.order
}
