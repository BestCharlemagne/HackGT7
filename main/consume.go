package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func consume() {
	response, err := http.Get("http://localhost:8085/articles")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))
}
