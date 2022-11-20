package handlers

import (
	"IS_Lab/src/model"
	"IS_Lab/src/parsers"
	"IS_Lab/src/store"
	"encoding/json"
	"log"
	"net/http"
)

type CarIndex struct {
	l *log.Logger
	s *store.Store
}

func NewCarIndexHandler(l *log.Logger, s *store.Store) *CarIndex {
	return &CarIndex{l, s}
}
func (h *CarIndex) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if !HandleBasicAuth(r, h.s, 3) {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("username or password are incorrect"))
		return
	}
	switch r.Method {
	case http.MethodGet:
		h.ReadMonthly(rw, r)
	case http.MethodPost:
		h.CreateMonthly(rw, r)
	default:
		h.l.Println(http.StatusNotFound)
	}

}

func (h *CarIndex) CreateMonthly(rw http.ResponseWriter, r *http.Request) {

	a := model.Period{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&a); err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	nums := parsers.ParsePeriod(a.Date)

	if err := h.s.CarIndex().Create(nums[0], nums[1]); err != nil {
		http.Error(rw, err.Error(), http.StatusConflict)
		return
	}
}
func (h *CarIndex) ReadMonthly(rw http.ResponseWriter, r *http.Request) {
	type PeriodType struct {
		Date string `json:"period"`
		Type bool   `json:"type"`
	}
	a := PeriodType{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&a); err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	nums := parsers.ParsePeriod(a.Date)

	var err error
	if a.Type {
		var mod *[]model.TypeIndex

		if len(nums) == 2 {
			mod, err = h.s.CarIndex().ReadType(nums[0], nums[1])
		} else {
			mod, err = h.s.CarIndex().ReadTypePeriod(nums)
		}
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}
		for _, ele := range *mod {
			encoder := json.NewEncoder(rw)
			if err = encoder.Encode(&ele); err != nil {
				http.Error(rw, err.Error(), http.StatusBadRequest)
				return
			}
		}
	} else {
		var mod *[]model.BrandIndex

		if len(nums) == 2 {
			mod, err = h.s.CarIndex().ReadBrand(nums[0], nums[1])
		} else {
			mod, err = h.s.CarIndex().ReadBrandPeriod(nums)
		}
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}
		for _, ele := range *mod {
			encoder := json.NewEncoder(rw)
			if err = encoder.Encode(&ele); err != nil {
				http.Error(rw, err.Error(), http.StatusBadRequest)
				return
			}
		}
	}

}
