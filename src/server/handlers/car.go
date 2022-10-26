package handlers

import (
	"IS_Lab/src/model"
	"IS_Lab/src/parsers"
	"IS_Lab/src/store"
	"encoding/json"
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
		h.GetCar(rw, r)

	default:
		h.l.Println(http.StatusNotFound)
	}

}

func (h *Car) CreateCar(rw http.ResponseWriter, r *http.Request) {
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
func (h *Car) GetCar(rw http.ResponseWriter, r *http.Request) {
	ints := parsers.GetIntsFromURL(r.URL.Path)
	if len(ints) == 0 {
		http.Error(rw, "Id was not provided", http.StatusBadRequest)
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
