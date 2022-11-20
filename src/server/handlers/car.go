package handlers

import (
	"IS_Lab/src/model"
	"IS_Lab/src/parsers"
	"IS_Lab/src/store"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Car struct {
	l *log.Logger
	s *store.Store
}

func NewCarHandler(l *log.Logger, s *store.Store) *Car {
	return &Car{l, s}
}
func (h *Car) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodGet:
		if !HandleBasicAuth(r, h.s, 1) {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte("username or password are incorrect"))
			return
		}
		h.ReadCar(rw, r)
	case http.MethodPost:
		if !HandleBasicAuth(r, h.s, 3) {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte("username or password are incorrect"))
			return
		}
		h.UpdateCar(rw, r)
	default:
		h.l.Println(http.StatusNotFound)
	}

}

func (h *Car) CreateCar(rw http.ResponseWriter, r *http.Request) {
	if !HandleBasicAuth(r, h.s, 3) {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("username or password are incorrect"))
		return
	}
	if r.Method != http.MethodPost {
		return
	}
	a := model.Car{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&a); err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.s.Car().Create(&a); err != nil {
		http.Error(rw, err.Error(), http.StatusConflict)
		http.Error(rw, "asd2", http.StatusBadRequest)
		return
	}
}
func (h *Car) ReadCar(rw http.ResponseWriter, r *http.Request) {
	ints := parsers.GetIntsFromURL(r.URL.Path)
	if len(ints) == 0 {
		fmt.Println("im here")
		a := model.CarChars{}
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&a); err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}
		cars, err := h.s.Car().Search(&a)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Println("im here 2")
		encoder := json.NewEncoder(rw)
		for i := range cars {
			if err = encoder.Encode(&cars[i]); err != nil {
				http.Error(rw, err.Error(), http.StatusBadGateway)
				return
			}
		}
		fmt.Println("im here 3")
		return
	}
	mod, err := h.s.Car().Read(ints[0])
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	encoder := json.NewEncoder(rw)
	if err = encoder.Encode(&mod); err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
}
func (h *Car) UpdateCar(rw http.ResponseWriter, r *http.Request) {
	ints := parsers.GetIntsFromURL(r.URL.Path)
	if len(ints) == 0 {
		http.Error(rw, "Id was not provided", http.StatusBadRequest)
	}
	type tempCar struct {
		Brand    string         `json:"brand,omitempty"`
		Model    string         `json:"model,omitempty"`
		Power    int            `json:"power,omitempty"`
		Capacity int            `json:"capacity,omitempty"`
		Fuel     model.FuelType `json:"fuel,omitempty"`
		Type     model.CarType  `json:"type,omitempty"`
		Price    int            `json:"price,omitempty"`
	}
	a := tempCar{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&a); err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	mod := model.Car{
		Id:       ints[0],
		Brand:    a.Brand,
		Model:    a.Model,
		Power:    a.Power,
		Capacity: a.Capacity,
		Fuel:     a.Fuel,
		Type:     a.Type,
		Price:    a.Price,
		IsSold:   false,
	}
	if err := h.s.Car().Update(&mod); err != nil {
		http.Error(rw, err.Error(), http.StatusConflict)
		return
	}
}
