package main

import (
	"fmt"
	"log"
	"net/http"
	"telesp/handlers"
)

type Page struct { // Declare page params
	Title string
	Body []byte
}


func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Don Anton Gun %s!", r.URL.Path[1:])
}

func main() {
    http.HandleFunc("/", handler)
	handlers.HandlersLunch("/") // This page insert pref str, handler instead

    log.Fatal(http.ListenAndServe(":8080", nil))
}