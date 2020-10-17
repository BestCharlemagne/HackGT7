package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"strconv"
)

type Store struct {
	ID int `json:"ID"`
	Title string `json:"Title"`
	Location int `json:"location"`
	Food string `json:"food"`
}

type Stores []Store
var stores Stores

func initStores() {
	stores = Stores {
		Store{ID: 9999, Title:"Joe's Market", Location: 30332, Food:"Apples, Bananas, and Pears"},
		Store{ID: 9887, Title:"Sam's Market", Location: 30482, Food:"Oranges, Bananas, and Pineapples"},
	}
}

func GetStoreFromID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: All Articles Endpoint")
	id := r.URL.Query().Get("ID")

	userId, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("Error converting string to int: %s", err)
		return
	}

	for  i := 0; i < len(stores); i++ {
		if stores[i].ID == userId {
			json.NewEncoder(w).Encode(stores[i])
			return
		}
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/stores", GetStoreFromID)
	//http.HandleFunc("/stores/{id}")
	log.Fatal(http.ListenAndServe(":8085", nil))
}

func main() {
	initStores()
	handleRequests()
}