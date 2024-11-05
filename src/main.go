package main

import (
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"telesp/src/web"
)

// type application struct {

// }

func main() {
	mux := http.NewServeMux()
	//mux.HandleFunc("/", http.HandlerFunc(web.Home)) // Превратили функцию home в некоторый возможный обрабочик http запросов
	mux.HandleFunc("/", http.HandlerFunc(web.IndexHandler))
	mux.HandleFunc("/send", web.SendHandler)
	// mux.HandleFunc("/snippet", web.ShowSnippet)
	// mux.HandleFunc("/snippet/create", web.CreateSnippet)

	fileServer := http.FileServer(http.Dir("./internal/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Запуск веб-сервера на http://127.0.0.1:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
