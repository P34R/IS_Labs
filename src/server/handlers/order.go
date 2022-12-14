package handlers

import (
	"IS_Lab/src/model"
	"IS_Lab/src/parsers"
	"IS_Lab/src/store"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Order struct {
	l *log.Logger
	s *store.Store
}

func NewOrderHandler(l *log.Logger, s *store.Store) *Order {
	return &Order{l, s}
}
func (h *Order) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodGet:
		if !HandleBasicAuth(r, h.s, 3) {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte("username or password are incorrect"))
			return
		}
		h.ReadOrder(rw, r)
	case http.MethodPost:
		if !HandleBasicAuth(r, h.s, 3) {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte("username or password are incorrect"))
			return
		}
		h.UpdateOrder(rw, r)
	default:
		h.l.Println(http.StatusNotFound)
	}

}

func (h *Order) CreateOrder(rw http.ResponseWriter, r *http.Request) {
	if !HandleBasicAuth(r, h.s, 1) {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("username or password are incorrect"))
		return
	}
	if r.Method != http.MethodPost {
		return
	}
	a := model.Order{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&a); err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	if a.Date == "" {
		y, m, d := time.Now().Date()
		a.Date = m.String() + "." + strconv.Itoa(d) + "." + strconv.Itoa(y)
	}
	if err := h.s.Order().Create(&a); err != nil {
		http.Error(rw, err.Error(), http.StatusConflict)
		return
	}
}
func (h *Order) ReadOrder(rw http.ResponseWriter, r *http.Request) {
	ints := parsers.GetIntsFromURL(r.URL.Path)
	if len(ints) == 0 {
		http.Error(rw, "Id was not provided", http.StatusBadRequest)
		return
	}
	mod, err := h.s.Order().Read(ints[0])
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

func (h *Order) UpdateOrder(rw http.ResponseWriter, r *http.Request) {
	ints := parsers.GetIntsFromURL(r.URL.Path)
	if len(ints) == 0 {
		http.Error(rw, "Id was not provided", http.StatusBadRequest)
	}
	type stat struct {
		Status model.Status `json:"status"`
	}
	var st stat
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&st); err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.s.Order().Update(ints[0], string(st.Status)); err != nil {
		http.Error(rw, err.Error(), http.StatusConflict)
		return
	}
}
