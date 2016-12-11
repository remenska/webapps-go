package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayhellowName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //parse arguments, you have to call this yourself
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])

	for k, v := range r.Form {
		fmt.Println("key: ", k)
		fmt.Println("value:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello there!") // send data to the client side
}

func main() {
	http.HandleFunc("/", sayhellowName)
	err := http.ListenAndServe(":9090", nil) // set the listen port
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
