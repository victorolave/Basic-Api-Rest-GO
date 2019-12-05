package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type Person struct {
	 ID string `json:"id, ommitempty"`
	 FirstName string `json:"firstname, ommitempty"`
	 LastName string `json:"lastname, ommitempty"`
	 Address Address `json:"address, ommitempty"`
}

type Address struct {
	City string `json:"city, ommitempty"`
	State string `json: "state, ommitempty"`
}

var people []Person

func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request){
	json.NewEncoder(w).Encode(people)
}

func GetPersonEndPoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&Person{})
}

func CreatePersonEndPoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var person Person
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

func DeletePersonEndPoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index + 1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(people)
}

func main() {
	router := mux.NewRouter()

	people = append(people, Person{ID:"1", FirstName:"Victor", LastName:"Olave", Address: Address{City:"Popayán", State:"Cauca"}})
	people = append(people, Person{ID:"2", FirstName:"Sandra", LastName:"Olave", Address: Address{City:"Popayán", State:"Cauca"}})
	people = append(people, Person{ID:"3", FirstName:"Gabriela", LastName:"Sarria", Address: Address{City:"Popayán", State:"Cauca"}})

	//Endpoints
	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetPersonEndPoint).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePersonEndPoint).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePersonEndPoint).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", router))
}