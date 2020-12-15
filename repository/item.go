package repository

//ItemType is an alias for an unsigned integer to represent different types of items
type ItemType uint

//Enum-like const structure
const (
	Product ItemType = iota
	Empty
	Wall
	Entrance
	Exit
)

func (i ItemType) String() string {
	return []string{"Product", "Empty", "Wall", "Entrance", "Exit"}[i]
}

//Item contains a format for an item to be sold at a store.
//Has the row and column that the item can be found at in the store grid.
//Also contains a type. This defaults product if not specified upon initialization
type Item struct {
	Type   ItemType
	Name   string
	Row    int
	Column int
}
