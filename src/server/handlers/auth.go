package handlers

import (
	"IS_Lab/src/store"
	"fmt"
	"net/http"
)

// 1 - any, 2 - user, 3 - admin
func HandleBasicAuth(r *http.Request, s *store.Store, code int) bool {
	if code < 1 || code > 3 {
		return false
	}
	log, pass, ok := r.BasicAuth()
	fmt.Println(log, pass, ok)
	if ok {
		u, err := s.User().Read(log)
		if err != nil {
			fmt.Println(err)
			fmt.Println("1")
			return false
		}
		err, b := u.CompareWithHash(pass)
		if b == false || err != nil {
			fmt.Println(b, err, "2")
			return b
		}
		u.Sanitize()
		if code == 2 && u.UserType != "user" {
			fmt.Println("3")
			return false
		} else if code == 3 && u.UserType != "admin" {
			fmt.Println("4")
			return false
		}
		return b
	}
	return false
}
