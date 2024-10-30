package main

import (
	"log"
	"net/http"
	"telesp/src/web"
)


func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", http.HandlerFunc(web.Home)) // Превратили функцию home в некоторый возможный обрабочик http запросов
	// mux.HandleFunc("/snippet", web.ShowSnippet)
	// mux.HandleFunc("/snippet/create", web.CreateSnippet)

	fileServer := http.FileServer(http.Dir("./internal/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
 
	log.Println("Запуск веб-сервера на http://127.0.0.1:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}