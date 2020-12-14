package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/BestCharlemagne/StoreNavigator/repository"
)

//GetStores takes in a splice of stores and returns either a list of stores back containing the
//desired output from the response writer or an error
func GetStores(stores []repository.Store, w http.ResponseWriter, r *http.Request) ([]repository.Store, error) {
	if id := r.URL.Query().Get("id"); id != "" {
		userID, err := strconv.Atoi(id)
		if err != nil {
			return nil, fmt.Errorf("Error converting string to int: %s", err)
		}

		for i := 0; i < len(stores); i++ {
			if stores[i].ID == userID {
				return repository.Stores{stores[i]}, nil
			}
		}
	} else if zip := r.URL.Query().Get("zipCode"); zip != "" {
		zipCode, err := strconv.Atoi(zip)
		if err != nil {
			return nil, fmt.Errorf("Error converting string to int: %s", err)
		}
		var copiedStores repository.Stores = make(repository.Stores, len(stores))
		copy(copiedStores, stores)

		quickSelect(zipCode, copiedStores, 0, len(copiedStores)-1, 5)
		return copiedStores[0:5], nil
	}
	return nil, errors.New("Error: Must provide a zipcode or id: try '/stores?zipCode=30332'")
}
