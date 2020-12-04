package main

type PathType [][]string

//Store contains data pertaining to a Store.
type Store struct {
	ID      int      `json:"id"`
	Title   string   `json:"title"`
	ZipCode int      `json:"zipCode"`
	Items   []Item   `json:"items"`
	Path    PathType `json:"path"`
}

//Stores is a type defintion for a slice of stores
type Stores []Store
