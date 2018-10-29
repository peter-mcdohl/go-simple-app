package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

// Display all from the people var
func GetPeople(w http.ResponseWriter, r *http.Request) {
	for _, item := range people {
		// if item.ID == params["id"] {
		json.NewEncoder(w).Encode(item)
		// return
		// }
	}
	json.NewEncoder(w).Encode(&Person{})
}

// Display a single data
func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, r.URL.Path)
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	router := mux.NewRouter()
	people = append(people, Person{ID: "1", Firstname: "paul", Lastname: "Logan", Address: &Address{City: "Bogor", State: "Indonesia"}})
	people = append(people, Person{ID: "2", Firstname: "kang", Lastname: "didi", Address: &Address{City: "Bandung", State: "cibereum"}})
	router.HandleFunc("/", handler)
	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
