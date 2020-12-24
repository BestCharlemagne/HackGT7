package repository

//PathType is used to decode and encode a json array of arrays into this Go type
type PathType [][]string

//ItemType is used to decode and encode a json arrays into this Go type
type ItemType []Item

//Store contains data pertaining to a Store.
type Store struct {
	ID      int      `json:"id"`
	Title   string   `json:"title"`
	ZipCode int      `json:"zipCode"`
	Items   ItemType `json:"items"`
	Path    PathType `json:"path"`
}

//Stores is a type defintion for a slice of stores
type Stores []Store
