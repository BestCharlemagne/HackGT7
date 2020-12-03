package main

import (
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
