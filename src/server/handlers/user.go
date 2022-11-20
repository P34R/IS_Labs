package handlers

import (
	"IS_Lab/src/model"
	"IS_Lab/src/store"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	l *log.Logger
	s *store.Store
}

func NewUserHandler(l *log.Logger, s *store.Store) *User {
	return &User{l, s}
}
func (h *User) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	if !HandleBasicAuth(r, h.s, 3) {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("username or password are incorrect"))
		return
	}

	switch r.Method {

	case http.MethodPost:
		h.CreateUser(rw, r)

	case http.MethodDelete:
		h.DeleteUser(rw, r)

	default:
		h.l.Println(http.StatusNotFound)
	}

}

func (h *User) CreateUser(rw http.ResponseWriter, r *http.Request) {
	a := model.User{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&a); err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(a)
	if err := h.s.User().Create(&a); err != nil {
		http.Error(rw, err.Error(), http.StatusConflict)
		return
	}
}

func (h *User) DeleteUser(rw http.ResponseWriter, r *http.Request) {
	var username string
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&username); err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.s.User().Delete(username); err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
}
