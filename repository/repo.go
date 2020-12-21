package repository

//Repo defines interface for a database repository
type Repo interface {
	AddStore(store Store)
	AddAllStores(stores []Store)
	RemoveStore(store Store)
	RemoveAllStores()
	GetAllStores() []Store
	Close()
}

//ExampleStores is an array of example stores of course!
var ExampleStores []Store = []Store{
	{
		Title:   "Joe's Market",
		ZipCode: 30332,
		Items: []Item{
			{Type: Product, ID: 0, Name: "Oranges", Row: 0, Column: 0},
			{Type: Product, ID: 1, Name: "Strawberry", Row: 3, Column: 6},
			{Type: Product, ID: 2, Name: "Banana", Row: 2, Column: 2}},
		Path: [][]pathRune{
			{'+', '+', '+', '+', '+', '+', '+'},
			{'+', '/', '+', '/', '+', '/', '+'},
			{'+', '/', '+', '/', '+', '/', '+'},
			{'+', '/', '+', '/', '+', '/', '+'},
			{'+', '/', '+', '/', '+', '/', '+'},
			{'+', '+', '+', '+', '+', '+', '+'},
			{'-', '+', '+', '+', '+', '+', '_'}}},
	{
		Title:   "Sam's Market",
		ZipCode: 30402,
		Items: []Item{
			{Type: Product, Name: "Oranges", Row: 0, Column: 2},
			{Type: Product, Name: "Banana", Row: 3, Column: 6},
			{Type: Product, Name: "Pineapple", Row: 5, Column: 3}},
		Path: [][]pathRune{
			{'+', '+', '+', '+', '+', '+', '+'},
			{'+', '/', '/', '/', '/', '/', '+'},
			{'+', '+', '+', '+', '+', '+', '+'},
			{'+', '+', '+', '+', '+', '+', '+'},
			{'+', '/', '/', '/', '/', '/', '+'},
			{'+', '+', '+', '+', '+', '+', '+'},
			{'-', '+', '+', '+', '+', '+', '_'}}},
	{
		Title:   "Rons's Market",
		ZipCode: 30082,
		Items: []Item{
			{Type: Product, Name: "Oranges", Row: 0, Column: 2},
			{Type: Product, Name: "Banana", Row: 3, Column: 6},
			{Type: Product, Name: "Pineapple", Row: 5, Column: 3}},
		Path: [][]pathRune{
			{'+', '+', '+', '+', '+', '+', '+'},
			{'+', '/', '/', '/', '/', '/', '+'},
			{'+', '+', '+', '+', '+', '+', '+'},
			{'+', '+', '+', '+', '+', '+', '+'},
			{'+', '/', '/', '/', '/', '/', '+'},
			{'+', '+', '+', '+', '+', '+', '+'},
			{'-', '+', '+', '+', '+', '+', '_'}}},
	{
		Title:   "Jessica's Market",
		ZipCode: 50482,
		Items: []Item{
			{Type: Product, Name: "Oranges", Row: 0, Column: 2},
			{Type: Product, Name: "Banana", Row: 3, Column: 6},
			{Type: Product, Name: "Pineapple", Row: 5, Column: 3}},
		Path: [][]pathRune{
			{'+', '+', '+', '+', '+', '+', '+'},
			{'+', '/', '/', '/', '/', '/', '+'},
			{'+', '+', '+', '+', '+', '+', '+'},
			{'+', '+', '+', '+', '+', '+', '+'},
			{'+', '/', '/', '/', '/', '/', '+'},
			{'+', '+', '+', '+', '+', '+', '+'},
			{'-', '+', '+', '+', '+', '+', '_'}}},
	{
		Title:   "Jessica's Market",
		ZipCode: 92117,
		Items: []Item{
			{Type: Product, Name: "Oranges", Row: 0, Column: 2},
			{Type: Product, Name: "Banana", Row: 3, Column: 6},
			{Type: Product, Name: "Pineapple", Row: 5, Column: 3}},
		Path: [][]pathRune{
			{'+', '+', '+', '+', '+', '+', '+'},
			{'+', '/', '/', '/', '/', '/', '+'},
			{'+', '+', '+', '+', '+', '+', '+'},
			{'+', '+', '+', '+', '+', '+', '+'},
			{'+', '/', '/', '/', '/', '/', '+'},
			{'+', '+', '+', '+', '+', '+', '+'},
			{'-', '+', '+', '+', '+', '+', '_'}}},
	{
		Title:   "Jessica's Market",
		ZipCode: 30040,
		Items: []Item{
			{Type: Product, Name: "Oranges", Row: 0, Column: 2},
			{Type: Product, Name: "Banana", Row: 3, Column: 6},
			{Type: Product, Name: "Pineapple", Row: 5, Column: 3}},
		Path: [][]pathRune{
			{'+', '+', '+', '+', '+', '+', '+'},
			{'+', '/', '/', '/', '/', '/', '+'},
			{'+', '+', '+', '+', '+', '+', '+'},
			{'+', '+', '+', '+', '+', '+', '+'},
			{'+', '/', '/', '/', '/', '/', '+'},
			{'+', '+', '+', '+', '+', '+', '+'},
			{'-', '+', '+', '+', '+', '+', '_'}}},
	{
		Title:   "Bobs's Market",
		ZipCode: 29982,
		Items: []Item{
			{Type: Product, Name: "Oranges", Row: 0, Column: 2},
			{Type: Product, Name: "Banana", Row: 3, Column: 6},
			{Type: Product, Name: "Pineapple", Row: 5, Column: 3}},
		Path: [][]pathRune{
			{'+', '+', '+', '+', '+', '+', '+'},
			{'+', '/', '/', '/', '/', '/', '+'},
			{'+', '+', '+', '+', '+', '+', '+'},
			{'+', '+', '+', '+', '+', '+', '+'},
			{'+', '/', '/', '/', '/', '/', '+'},
			{'+', '+', '+', '+', '+', '+', '+'},
			{'-', '+', '+', '+', '+', '+', '_'}}},
	{
		Title:   "Liam's Market",
		ZipCode: 40082,
		Items: []Item{
			{Type: Product, Name: "Oranges", Row: 0, Column: 2},
			{Type: Product, Name: "Banana", Row: 3, Column: 6},
			{Type: Product, Name: "Pineapple", Row: 5, Column: 3}},
		Path: [][]pathRune{
			{'+', '+', '+', '+', '+', '+', '+'},
			{'+', '/', '/', '/', '/', '/', '+'},
			{'+', '+', '+', '+', '+', '+', '+'},
			{'+', '+', '+', '+', '+', '+', '+'},
			{'+', '/', '/', '/', '/', '/', '+'},
			{'+', '+', '+', '+', '+', '+', '+'},
			{'-', '+', '+', '+', '+', '+', '_'}}},
	{
		Title:   "Jessica's Market",
		ZipCode: 20482,
		Items: []Item{
			{Type: Product, Name: "Oranges", Row: 0, Column: 2},
			{Type: Product, Name: "Banana", Row: 3, Column: 6},
			{Type: Product, Name: "Pineapple", Row: 5, Column: 3}},
		Path: [][]pathRune{
			{'+', '+', '+', '+', '+', '+', '+'},
			{'+', '/', '/', '/', '/', '/', '+'},
			{'+', '+', '+', '+', '+', '+', '+'},
			{'+', '+', '+', '+', '+', '+', '+'},
			{'+', '/', '/', '/', '/', '/', '+'},
			{'+', '+', '+', '+', '+', '+', '+'},
			{'-', '+', '+', '+', '+', '+', '_'}}},
	{
		Title:   "Jessica's Market",
		ZipCode: 30002,
		Items: []Item{
			{Type: Product, Name: "Oranges", Row: 0, Column: 2},
			{Type: Product, Name: "Banana", Row: 3, Column: 6},
			{Type: Product, Name: "Pineapple", Row: 5, Column: 3}},
		Path: [][]pathRune{
			{'+', '+', '+', '+', '+', '+', '+'},
			{'+', '/', '/', '/', '/', '/', '+'},
			{'+', '+', '+', '+', '+', '+', '+'},
			{'+', '+', '+', '+', '+', '+', '+'},
			{'+', '/', '/', '/', '/', '/', '+'},
			{'+', '+', '+', '+', '+', '+', '+'},
			{'-', '+', '+', '+', '+', '+', '_'}}}}
