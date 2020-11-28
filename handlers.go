package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//GetStore returns all store ids closes to provided zip code.
//If an id is presented, then will return a Store struct.
func GetStore(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: All Articles Endpoint")
	stores, err := GetStores(stores, w, r)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	} else if stores != nil && len(stores) > 0 {
		json.NewEncoder(w).Encode(stores)
	} else {
		json.NewEncoder(w).Encode("Could not find any stores")
	}
}

//CalculatePath expects a query string with a list of items to get.
//Will return an array of the path to take between the items.
func CalculatePath(w http.ResponseWriter, r *http.Request) {

}

//HomePage just prints  simple message saying you are at the homepage.
func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Api to get store paths!")
}
