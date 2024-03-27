package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/johannasiagian/go_students_crud_mysql/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterStudentsRoutes(r)
	http.Handle("/", r)
	fmt.Print("Starting Server localhost:9010")
	log.Fatal(http.ListenAndServe("Localhost:9010", r))
}
