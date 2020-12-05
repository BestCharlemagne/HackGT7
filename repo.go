package main

import (
	"fmt"
)

var currentID int

var stores Stores

//Repo defines interface for a database repository
type Repo interface {
	AddStore(store Store)
	AddAllStores(stores []Store)
	// RemoveStores(store Store)
	GetStores() []Store
	Close()
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
