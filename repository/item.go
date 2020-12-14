package repository

//Item contains a format for an item to be sold at a store.
//Has the row and column that the item can be found at in the store grid.
type Item struct {
	Name   string
	Row    int
	Column int
}
