package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name"`
	City    string `json:"city"`
	Zipcode string `json:"zip_code"`
}

func main() {
	// define  routes
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getAllCustomers)
	// start the server
	http.ListenAndServe("localhost:8000", nil)
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{Name: "John Doe", City: "San Francisco", Zipcode: "94105"},
		{Name: "Jane Doe", City: "San Francisco", Zipcode: "94105"},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)

}
