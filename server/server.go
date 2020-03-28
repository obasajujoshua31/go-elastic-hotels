package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

type Router struct {
	*mux.Router
}

func NewRouter() *Router {
	return &Router{mux.NewRouter()}
}

func (r *Router) initializeRoutes() {
	r.HandleFunc("/", GetHomeHandler()).Methods(http.MethodGet)
	r.HandleFunc("/hotels/{name}", SearchHotelHandler()).Methods(http.MethodGet)
}

func Start() error {
	r := NewRouter()
	r.initializeRoutes()
	log.Printf("Server starting at port :%s", os.Getenv("APP_PORT"))
	if err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("APP_PORT")), r); err != nil {
		return err
	}
	return nil
}
