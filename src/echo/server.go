package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func homePage(w http.ResponseWriter, r *http.Request){
	//fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
	for header, values := range r.Header {
		fmt.Printf( "%v: %v\n", header, strings.Join(values, " "))
		fmt.Fprintf(w, "%v: %v\n", header, strings.Join(values, " "))
	}
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":9000", nil))
}

