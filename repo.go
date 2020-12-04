package main

import (
	"encoding/json"
	"fmt"
	"log"
)

var currentID int

var stores Stores

type Repo interface {
	Open()
	Close()
	addStore(store Store)
	addAllStores(stores []Store)
	removeStores(store Store)
	getStores() []Store
}

func (m *PathType) Scan(src interface{}) error {
	// The data stored in a JSON field is actually returned as []uint8
	val := src.([]uint8)
	return json.Unmarshal(val, &m)
}

//GetAllStores currently retieves all stores with only support for a test SQL database.
func GetAllStores(test bool) {
	if test {
		repo := TestRepo{}
		database := repo.GetDatabase()

		rows, err := database.Query("SELECT * FROM stores")
		defer rows.Close()
		if err != nil {
			log.Fatal(err)
		} else {
			for rows.Next() {
				var newStore Store
				err := rows.Scan(&newStore.ID, &newStore.Title, &newStore.ZipCode, &newStore.Path)
				if err != nil {
					log.Fatal(err)
				}
				stores = append(stores, newStore)
			}
		}
		print(stores[0].Path[0][0])
	}
}

func PostData() {
	repo := TestRepo{}
	database := repo.GetDatabase()
	store := Store{
		Title:   "Joe's Market",
		ZipCode: 30332,
		Items: []Item{
			{Name: "Oranges", Row: 0, Column: 0},
			{Name: "Strawberry", Row: 3, Column: 6},
			{Name: "Banana", Row: 2, Column: 2}},
		Path: [][]string{
			{"+", "+", "+", "+", "+", "+", "+"},
			{"+", "/", "+", "/", "+", "/", "+"},
			{"+", "/", "+", "/", "+", "/", "+"},
			{"+", "/", "+", "/", "+", "/", "+"},
			{"+", "/", "+", "/", "+", "/", "+"},
			{"+", "+", "+", "+", "+", "+", "+"}}}
	bytes, err := json.Marshal(store.Path)
	if err == nil {
		stmt, err := database.Prepare("INSERT INTO stores VALUES (1, 'Store 1', 92117, ?)")
		if err != nil {
			log.Fatal(err)
		}
		result, err2 := stmt.Exec(bytes)

		if err2 == nil {
			print(result.RowsAffected())
		} else {
			log.Fatal(err2)
		}
	} else {
		log.Fatal(err)
	}
}

//RepoAddStores takes a slice of stores and appends them to the database.
func RepoAddStores(storesToAdd []Store) {
	for _, store := range storesToAdd {
		store.ID = currentID
		currentID++
	}
	stores = append(stores, storesToAdd...)
}

//RepoAddStore adds a store to the database.
func RepoAddStore(store Store) Store {
	store.ID = currentID
	currentID++
	stores = append(stores, store)
	return store
}

//RepoDestroyStore removes a store from the database of given id.
func RepoDestroyStore(id int) error {
	for i, t := range stores {
		if t.ID == id {
			stores = append(stores[:i], stores[i+1:]...)
		}
	}
	return fmt.Errorf("Could not find Store with id of %d to delete", id)
}
