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
	mux.HandleFunc("/add", web.AddHandler)
	// mux.HandleFunc("/snippet", web.ShowSnippet)
	// mux.HandleFunc("/snippet/create", web.CreateSnippet)

	fileServer := http.FileServer(http.Dir("./internal/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Запуск веб-сервера на http://127.0.0.1:8080")
	err := http.ListenAndServe("localhost:8080", mux)
	log.Fatal(err)
}
