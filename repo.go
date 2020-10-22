package main

import "fmt"

var currentID int

var stores Stores

func init() {
	RepoCreateStore(
		Store{
			Title:   "Joe's Market",
			ZipCode: 30332,
			Items: []Item{
				Item{Name: "Oranges", Row: 0, Column: 0},
				Item{Name: "Strawberry", Row: 3, Column: 6},
				Item{Name: "Banana", Row: 2, Column: 2}},
			Path: [][]string{
				{"+", "+", "+", "+", "+", "+", "+"},
				{"+", "/", "+", "/", "+", "/", "+"},
				{"+", "/", "+", "/", "+", "/", "+"},
				{"+", "/", "+", "/", "+", "/", "+"},
				{"+", "/", "+", "/", "+", "/", "+"},
				{"+", "+", "+", "+", "+", "+", "+"}}})
	RepoCreateStore(
		Store{
			Title:   "Sam's Market",
			ZipCode: 30482,
			Items: []Item{
				Item{Name: "Oranges", Row: 0, Column: 2},
				Item{Name: "Banana", Row: 3, Column: 6},
				Item{Name: "Pineapple", Row: 5, Column: 3}},
			Path: [][]string{
				{"+", "+", "+", "+", "+", "+", "+"},
				{"+", "/", "/", "/", "/", "/", "+"},
				{"+", "+", "+", "+", "+", "+", "+"},
				{"+", "+", "+", "+", "+", "+", "+"},
				{"+", "/", "/", "/", "/", "/", "+"},
				{"+", "+", "+", "+", "+", "+", "+"}}})
	RepoCreateStore(
		Store{
			Title:   "Rons's Market",
			ZipCode: 30482,
			Items: []Item{
				Item{Name: "Oranges", Row: 0, Column: 2},
				Item{Name: "Banana", Row: 3, Column: 6},
				Item{Name: "Pineapple", Row: 5, Column: 3}},
			Path: [][]string{
				{"+", "+", "+", "+", "+", "+", "+"},
				{"+", "/", "/", "/", "/", "/", "+"},
				{"+", "+", "+", "+", "+", "+", "+"},
				{"+", "+", "+", "+", "+", "+", "+"},
				{"+", "/", "/", "/", "/", "/", "+"},
				{"+", "+", "+", "+", "+", "+", "+"}}})
	RepoCreateStore(
		Store{
			Title:   "Bobs's Market",
			ZipCode: 30482,
			Items: []Item{
				Item{Name: "Oranges", Row: 0, Column: 2},
				Item{Name: "Banana", Row: 3, Column: 6},
				Item{Name: "Pineapple", Row: 5, Column: 3}},
			Path: [][]string{
				{"+", "+", "+", "+", "+", "+", "+"},
				{"+", "/", "/", "/", "/", "/", "+"},
				{"+", "+", "+", "+", "+", "+", "+"},
				{"+", "+", "+", "+", "+", "+", "+"},
				{"+", "/", "/", "/", "/", "/", "+"},
				{"+", "+", "+", "+", "+", "+", "+"}}})
	RepoCreateStore(
		Store{
			Title:   "Liam's Market",
			ZipCode: 30482,
			Items: []Item{
				Item{Name: "Oranges", Row: 0, Column: 2},
				Item{Name: "Banana", Row: 3, Column: 6},
				Item{Name: "Pineapple", Row: 5, Column: 3}},
			Path: [][]string{
				{"+", "+", "+", "+", "+", "+", "+"},
				{"+", "/", "/", "/", "/", "/", "+"},
				{"+", "+", "+", "+", "+", "+", "+"},
				{"+", "+", "+", "+", "+", "+", "+"},
				{"+", "/", "/", "/", "/", "/", "+"},
				{"+", "+", "+", "+", "+", "+", "+"}}})
	RepoCreateStore(
		Store{
			Title:   "Jessica's Market",
			ZipCode: 30482,
			Items: []Item{
				Item{Name: "Oranges", Row: 0, Column: 2},
				Item{Name: "Banana", Row: 3, Column: 6},
				Item{Name: "Pineapple", Row: 5, Column: 3}},
			Path: [][]string{
				{"+", "+", "+", "+", "+", "+", "+"},
				{"+", "/", "/", "/", "/", "/", "+"},
				{"+", "+", "+", "+", "+", "+", "+"},
				{"+", "+", "+", "+", "+", "+", "+"},
				{"+", "/", "/", "/", "/", "/", "+"},
				{"+", "+", "+", "+", "+", "+", "+"}}})
}

//RepoCreateStore adds a store to the database.
func RepoCreateStore(store Store) Store {
	currentID++
	store.ID = currentID
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
