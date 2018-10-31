package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
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
	json.NewEncoder(w).Encode(people)
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

// create a new item
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

// update a new item
func EditPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	newPerson := new(Person)
	bodyBytes, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(bodyBytes, &newPerson)

	// newPerson := new(Person)
	for i, _ := range people {
		if people[i].ID == params["id"] {
			// fmt.Println(newPerson)
			// fmt.Println(newPerson.Firstname)
			fmt.Println(newPerson.Lastname)
			people[i] = *newPerson
			// people[i].Firstname = newPerson.Firstname
			// people[i].Lastname = newPerson.Lastname
			// people[i].Address.City = params["address"]["city"]
			// people[i].Address.State = params["address"]["state"]
			json.NewEncoder(w).Encode(people[i])
			return
		}
	}
}

// Delete an item
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(people)
	}
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
	router.HandleFunc("/people/{id}", EditPerson).Methods("PUT")
	router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
