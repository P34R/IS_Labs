package handlers

import (
	"IS_Lab/src/model"
	"IS_Lab/src/parsers"
	"IS_Lab/src/store"
	"encoding/json"
	"log"
	"net/http"
)

type Client struct {
	l *log.Logger
	s *store.Store
}

func NewClientHandler(l *log.Logger, s *store.Store) *Client {
	return &Client{l, s}
}
func (h *Client) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodGet:
		if !HandleBasicAuth(r, h.s, 3) {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte("username or password are incorrect"))
			return
		}
		h.ReadClient(rw, r)

	default:
		h.l.Println(http.StatusNotFound)
	}

}

func (h *Client) CreateClient(rw http.ResponseWriter, r *http.Request) {
	if !HandleBasicAuth(r, h.s, 1) {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("username or password are incorrect"))
		return
	}
	if r.Method != http.MethodPost {
		return
	}
	a := model.Client{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&a); err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.s.Client().Create(&a); err != nil {
		http.Error(rw, err.Error(), http.StatusConflict)
		return
	}
}
func (h *Client) ReadClient(rw http.ResponseWriter, r *http.Request) {
	ints := parsers.GetIntsFromURL(r.URL.Path)
	if len(ints) == 0 {
		http.Error(rw, "Id was not provided", http.StatusBadRequest)
		return
	}
	mod, err := h.s.Client().Read(ints[0])
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
