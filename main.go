package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	. "github.com/programadriano/go-restapi/config"
	. "github.com/programadriano/go-restapi/dao"
	. "github.com/programadriano/go-restapi/models"
)

var config = Config{}
var dao = MoviesDAO{}

func getAll(w http.ResponseWriter, r *http.Request) {
	movies, err := dao.getAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, movies)
}

func getByID(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func update(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func delete(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/movies", getAll).Methods("GET")
	r.HandleFunc("/api/v1/movies/{id}", getByID).Methods("GET")
	r.HandleFunc("/api/v1/movies", create).Methods("POST")
	r.HandleFunc("/api/v1/movies", update).Methods("PUT")
	r.HandleFunc("/api/v1/movies", delete).Methods("DELETE")

	var port = ":3000"
	fmt.Println("Server running in port:", port)
	log.Fatal(http.ListenAndServe(port, r))
}
