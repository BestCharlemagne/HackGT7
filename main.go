package main

import (
	"flag"
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
	test := flag.Bool("test", false, "Use to run on the test server. \nfalse by default.")
	flag.Parse()

	if *test {
		AddStores()
	}
	handleRequests()
}
