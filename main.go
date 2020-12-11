package main

import (
	"log"
	"net/http"
)

func handleRequests() {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/stores", GetStore)
	http.HandleFunc("/path", CalculatePath)
	log.Fatal(http.ListenAndServe(":8085", nil))
}

func main() {
	// test := flag.Bool("test", false, "Use to run on the test server. \nfalse by default.")
	// flag.Parse()

	testRepo := TestRepo{}

	testRepo.AddStore(Store{ID: 1, Title: "Trader Joe's", ZipCode: 30332, Items: []Item{
		{Name: "Oranges", Row: 0, Column: 2},
		{Name: "Banana", Row: 3, Column: 6},
		{Name: "Pineapple", Row: 5, Column: 3}},
		Path: [][]string{
			{"+", "+", "+", "+", "+", "+", "+"},
			{"+", "/", "/", "/", "/", "/", "+"},
			{"+", "+", "+", "+", "+", "+", "+"},
			{"+", "+", "+", "+", "+", "+", "+"},
			{"+", "/", "/", "/", "/", "/", "+"},
			{"+", "+", "+", "+", "+", "+", "+"}}})
	// stores := testRepo.GetAllStores()
	// print(stores)

	// testRepo.RemoveStore(Store{ID: 1})
	handleRequests()
}
