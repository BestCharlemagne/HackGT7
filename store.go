package main

//Store contains data pertaining to a Store.
type Store struct {
	ID      int        `json:"id"`
	Title   string     `json:"title"`
	ZipCode int        `json:"zipCode"`
	Items   []Item     `json:"items"`
	Path    [][]string `json:"path"`
}

func (s *Store) compareTo(other *Store) int {
	return s.ZipCode - other.ZipCode
}

//Stores is a type defintion for a slice of stores
type Stores []Store
