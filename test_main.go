package main

import "testing"

//AddStores adds testing store data to the repo
func AddStores() {
	RepoAddStores(
		[]Store{
			{
				Title:   "Joe's Market",
				ZipCode: 30332,
				Items: []Item{
					{Name: "Oranges", Row: 0, Column: 0},
					{Name: "Strawberry", Row: 3, Column: 6},
					{Name: "Banana", Row: 2, Column: 2}},
				Path: [][]string{
					{"+", "+", "+", "+", "+", "+", "+"},
					{"+", "/", "+", "/", "+", "/", "+"},
					{"+", "/", "+", "/", "+", "/", "+"},
					{"+", "/", "+", "/", "+", "/", "+"},
					{"+", "/", "+", "/", "+", "/", "+"},
					{"+", "+", "+", "+", "+", "+", "+"}}},
			{
				Title:   "Sam's Market",
				ZipCode: 30402,
				Items: []Item{
					{Name: "Oranges", Row: 0, Column: 2},
					{Name: "Banana", Row: 3, Column: 6},
					{Name: "Pineapple", Row: 5, Column: 3}},
				Path: [][]string{
					{"+", "+", "+", "+", "+", "+", "+"},
					{"+", "/", "/", "/", "/", "/", "+"},
					{"+", "+", "+", "+", "+", "+", "+"},
					{"+", "+", "+", "+", "+", "+", "+"},
					{"+", "/", "/", "/", "/", "/", "+"},
					{"+", "+", "+", "+", "+", "+", "+"}}},
			{
				Title:   "Rons's Market",
				ZipCode: 30082,
				Items: []Item{
					{Name: "Oranges", Row: 0, Column: 2},
					{Name: "Banana", Row: 3, Column: 6},
					{Name: "Pineapple", Row: 5, Column: 3}},
				Path: [][]string{
					{"+", "+", "+", "+", "+", "+", "+"},
					{"+", "/", "/", "/", "/", "/", "+"},
					{"+", "+", "+", "+", "+", "+", "+"},
					{"+", "+", "+", "+", "+", "+", "+"},
					{"+", "/", "/", "/", "/", "/", "+"},
					{"+", "+", "+", "+", "+", "+", "+"}}},
			{
				Title:   "Jessica's Market",
				ZipCode: 50482,
				Items: []Item{
					{Name: "Oranges", Row: 0, Column: 2},
					{Name: "Banana", Row: 3, Column: 6},
					{Name: "Pineapple", Row: 5, Column: 3}},
				Path: [][]string{
					{"+", "+", "+", "+", "+", "+", "+"},
					{"+", "/", "/", "/", "/", "/", "+"},
					{"+", "+", "+", "+", "+", "+", "+"},
					{"+", "+", "+", "+", "+", "+", "+"},
					{"+", "/", "/", "/", "/", "/", "+"},
					{"+", "+", "+", "+", "+", "+", "+"}}},
			{
				Title:   "Jessica's Market",
				ZipCode: 92117,
				Items: []Item{
					{Name: "Oranges", Row: 0, Column: 2},
					{Name: "Banana", Row: 3, Column: 6},
					{Name: "Pineapple", Row: 5, Column: 3}},
				Path: [][]string{
					{"+", "+", "+", "+", "+", "+", "+"},
					{"+", "/", "/", "/", "/", "/", "+"},
					{"+", "+", "+", "+", "+", "+", "+"},
					{"+", "+", "+", "+", "+", "+", "+"},
					{"+", "/", "/", "/", "/", "/", "+"},
					{"+", "+", "+", "+", "+", "+", "+"}}},
			{
				Title:   "Jessica's Market",
				ZipCode: 30040,
				Items: []Item{
					{Name: "Oranges", Row: 0, Column: 2},
					{Name: "Banana", Row: 3, Column: 6},
					{Name: "Pineapple", Row: 5, Column: 3}},
				Path: [][]string{
					{"+", "+", "+", "+", "+", "+", "+"},
					{"+", "/", "/", "/", "/", "/", "+"},
					{"+", "+", "+", "+", "+", "+", "+"},
					{"+", "+", "+", "+", "+", "+", "+"},
					{"+", "/", "/", "/", "/", "/", "+"},
					{"+", "+", "+", "+", "+", "+", "+"}}},
			{
				Title:   "Bobs's Market",
				ZipCode: 29982,
				Items: []Item{
					{Name: "Oranges", Row: 0, Column: 2},
					{Name: "Banana", Row: 3, Column: 6},
					{Name: "Pineapple", Row: 5, Column: 3}},
				Path: [][]string{
					{"+", "+", "+", "+", "+", "+", "+"},
					{"+", "/", "/", "/", "/", "/", "+"},
					{"+", "+", "+", "+", "+", "+", "+"},
					{"+", "+", "+", "+", "+", "+", "+"},
					{"+", "/", "/", "/", "/", "/", "+"},
					{"+", "+", "+", "+", "+", "+", "+"}}},
			{
				Title:   "Liam's Market",
				ZipCode: 40082,
				Items: []Item{
					{Name: "Oranges", Row: 0, Column: 2},
					{Name: "Banana", Row: 3, Column: 6},
					{Name: "Pineapple", Row: 5, Column: 3}},
				Path: [][]string{
					{"+", "+", "+", "+", "+", "+", "+"},
					{"+", "/", "/", "/", "/", "/", "+"},
					{"+", "+", "+", "+", "+", "+", "+"},
					{"+", "+", "+", "+", "+", "+", "+"},
					{"+", "/", "/", "/", "/", "/", "+"},
					{"+", "+", "+", "+", "+", "+", "+"}}},
			{
				Title:   "Jessica's Market",
				ZipCode: 20482,
				Items: []Item{
					{Name: "Oranges", Row: 0, Column: 2},
					{Name: "Banana", Row: 3, Column: 6},
					{Name: "Pineapple", Row: 5, Column: 3}},
				Path: [][]string{
					{"+", "+", "+", "+", "+", "+", "+"},
					{"+", "/", "/", "/", "/", "/", "+"},
					{"+", "+", "+", "+", "+", "+", "+"},
					{"+", "+", "+", "+", "+", "+", "+"},
					{"+", "/", "/", "/", "/", "/", "+"},
					{"+", "+", "+", "+", "+", "+", "+"}}},
			{
				Title:   "Jessica's Market",
				ZipCode: 30002,
				Items: []Item{
					{Name: "Oranges", Row: 0, Column: 2},
					{Name: "Banana", Row: 3, Column: 6},
					{Name: "Pineapple", Row: 5, Column: 3}},
				Path: [][]string{
					{"+", "+", "+", "+", "+", "+", "+"},
					{"+", "/", "/", "/", "/", "/", "+"},
					{"+", "+", "+", "+", "+", "+", "+"},
					{"+", "+", "+", "+", "+", "+", "+"},
					{"+", "/", "/", "/", "/", "/", "+"},
					{"+", "+", "+", "+", "+", "+", "+"}}},
		})
}

//TestMain tests the main function
func TestMain(t *testing.T) {
}
