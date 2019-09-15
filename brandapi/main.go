package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Brand : struct
type Brand struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Piece    string    `json:"piece"`
	Designer *Designer `json:"designer"`
}

// Designer : struct
type Designer struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func getBrands(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(brands)
}

//get one brand
func getBrand(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) // get parameters
	// loopy
	for _, item := range brands {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Brand{})
}

//create brand
func createBrand(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var brand Brand
	_ = json.NewDecoder(r.Body).Decode(&brand)
	brand.ID = strconv.Itoa(rand.Intn(10000000)) // mock id
	brands = append(brands, brand)
	json.NewEncoder(w).Encode(brand)
}

//update brand
func updateBrand(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range brands {
		if item.ID == params["id"] {
			brands = append(brands[:index], brands[index+1:]...)
			var brand Brand
			_ = json.NewDecoder(r.Body).Decode(&brand)
			brand.ID = params["id"] // mock id
			brands = append(brands, brand)
			json.NewEncoder(w).Encode(brand)
			return

		}
	}
	json.NewEncoder(w).Encode(brands)
}

//delete brand
func deleteBrand(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range brands {
		if item.ID == params["id"] {
			brands = append(brands[:index], brands[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(brands)
}

// init brand var as slice Brand struct
var brands []Brand

func main() {
	r := mux.NewRouter()

	//mock data
	brands = append(brands, Brand{ID: "1", Isbn: "1234", Piece: "Hoodie", Designer: &Designer{Firstname: "Supreme", Lastname: "New York"}})
	brands = append(brands, Brand{ID: "2", Isbn: "12345", Piece: "T shirt", Designer: &Designer{Firstname: "Helmut", Lastname: "Lang"}})

	// router handlers / endpoint
	r.HandleFunc("/api/brands", getBrands).Methods("GET")
	r.HandleFunc("/api/brands/{id}", getBrand).Methods("GET")
	r.HandleFunc("/api/brands", createBrand).Methods("POST")
	r.HandleFunc("/api/brands/{id}", updateBrand).Methods("PUT")
	r.HandleFunc("/api/brands/{id}", deleteBrand).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", r))
}
