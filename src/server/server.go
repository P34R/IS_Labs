package server

import (
	"IS_Lab/src/server/handlers"
	"IS_Lab/src/store"
	"log"
	"net/http"
	"os"
	"time"
)

type Server struct {
	store    *store.Store
	ServeMux *http.ServeMux
	Server   *http.Server
}

func newServer() *Server {
	ServeMux := http.NewServeMux()
	s := store.NewStore()
	return &Server{
		store:    s,
		ServeMux: ServeMux,
		Server: &http.Server{
			Addr:         ":9091",
			Handler:      ServeMux,
			IdleTimeout:  120 * time.Second,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 5 * time.Second,
		},
	}
}
func (s *Server) configure() {
	l := log.New(os.Stdout, "project-api ", log.LstdFlags)
	ch := handlers.NewCarHandler(l, s.store)
	clh := handlers.NewClientHandler(l, s.store)
	oh := handlers.NewOrderHandler(l, s.store)
	uh := handlers.NewUserHandler(l, s.store)
	cih := handlers.NewCarIndexHandler(l, s.store)

	s.ServeMux.Handle("/cars/", ch)
	s.ServeMux.Handle("/clients/", clh)
	s.ServeMux.Handle("/orders/", oh)
	s.ServeMux.Handle("/user", uh)
	s.ServeMux.Handle("/car_index", cih)
	s.ServeMux.HandleFunc("/car", ch.CreateCar)
	s.ServeMux.HandleFunc("/client", clh.CreateClient)
	s.ServeMux.HandleFunc("/order", oh.CreateOrder)

}
func Start() {
	s := newServer()
	s.configure()
	if err := s.store.Open(); err != nil {
		panic(err)
	}
	defer s.store.Close()
	if err := s.Server.ListenAndServe(); err != nil {
		panic(err)
	}
}
