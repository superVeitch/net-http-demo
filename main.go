package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"testing"
)

func main()  {
	http.HandleFunc("/foo", handleIndex)
	http.HandleFunc("/host", handleUp)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleIndex(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func handleUp(w http.ResponseWriter, r *http.Request) {
	if r.Host == "localhost" {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.Host))
	} else {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString("127.0.0.1"))
	}
}

//func dumpCover(w http.ResponseWriter, r *http.Request)  {
//	w.Write(testing.GetCover())
//	fmt.Fprintf(w, testing.GetCover())
//}