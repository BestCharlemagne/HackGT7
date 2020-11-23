package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

//GetStore returns all store ids closes to provided zip code.
//If an id is presented, then will return a Store struct.
func GetStore(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: All Articles Endpoint")

	if id := r.URL.Query().Get("id"); id != "" {
		userID, err := strconv.Atoi(id)
		if err != nil {
			log.Printf("Error converting string to int: %s", err)
			return
		}

		for i := 0; i < len(stores); i++ {
			if stores[i].ID == userID {
				json.NewEncoder(w).Encode(stores[i])
			}
		}
	} else if zip := r.URL.Query().Get("zipCode"); zip != "" {
		zipCode, err := strconv.Atoi(zip)
		if err != nil {
			log.Printf("Error converting string to int: %s", err)
		}

		json.NewEncoder(w).Encode(quickSelect(zipCode, stores, 0, len(stores)-1, 5))
	} else {
		json.NewEncoder(w).Encode("Error: Must provide a zipcode or id: try '/stores?zipCode=30332'")
	}
}

func quickSelect(zipCode int, array Stores, start int, end int, numItems int) Store {
	pivot := len(array) / 2
	pivotItem := (array)[pivot]
	swap(&array, 0, pivot) //Pivot now at front

	i := start
	j := end

	for i <= j {
		for i <= j && abs(zipCode-(array)[i].ZipCode)-pivotItem.ZipCode < 0 {
			i++
		}
		for i <= j && abs(zipCode-(array)[i].ZipCode)-pivotItem.ZipCode > 0 {
			j++
		}
		if i <= j {
			swap(&array, i, j)
			i++
			j--
		}
	}
	swap(&array, start, j) //Swap pivot into correct place
	if j == numItems-1 {
		return array[j]
	} else if j > numItems-1 {
		return quickSelect(zipCode, array, start, j-1, numItems)
	}
	return quickSelect(zipCode, array, j+1, end, numItems)
}

func swap(array *Stores, to int, from int) {
	store := (*array)[to]
	(*array)[to] = (*array)[from]
	(*array)[from] = store
}

//CalculatePath expects a query string with a list of items to get.
//Will return an array of the path to take between the items.
func CalculatePath(w http.ResponseWriter, r *http.Request) {

}

//HomePage just prints  simple message saying you are at the homepage.
func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Api to get store paths!")
}
