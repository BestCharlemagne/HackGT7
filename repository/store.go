package repository

//PathArray is used to decode and encode a json array of arrays into this Go type
type PathArray [][]pathRune

//ItemArray is used to decode and encode a json arrays into this Go type
type ItemArray []Item

type pathRune rune

//Store contains data pertaining to a Store.
type Store struct {
	ID      int       `json:"id"`
	Title   string    `json:"title"`
	ZipCode int       `json:"zipCode"`
	Items   ItemArray `json:"items"`
	Path    PathArray `json:"path"`
}

//Stores is a type defintion for a slice of stores
type Stores []Store

func (r pathRune) getItemType() ItemType {
	switch r {
	case '+':
		return Empty
	case '/':
		return Wall
	case '-':
		return Entrance
	case '_':
		return Exit
	default:
		return Empty
	}
}
