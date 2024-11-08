package web

import (
	"fmt"
	"log"
	"telesp/pkg/models"
	"telesp/pkg/models/psql"

	// "fmt"
	"html/template"
	"log/slog"
	"net/http"
	// "strconv"
)

const countOfParams = 1

type Message struct {
	Message [countOfParams]string `json:"message"`
}

var storageData = models.TestPerson{}
var logger = slog.Logger{}

func Home(w http.ResponseWriter, r *http.Request) {
	// Check if cur url match "/" template
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./internal/html/homePage.html",
		"./internal/html/BaseLayout.html",
		"./internal/html/footer.partial.html",
	}

	tp, err := template.ParseFiles(files...) // or use home.page.tmpl
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = tp.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}

	// TODO: Add fully call of anything SQL func

	// w.Write([]byte("Привет из Snippetbox"))
}

func SendHandler(w http.ResponseWriter, r *http.Request) {
	var msg Message
	log.Println("SendHandler function called")
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", 405)
	}

	// читаем параметр message из формы
	for i := 0; i < countOfParams; i++ {
		msg.Message[i] = r.FormValue("field" + string(i+1))
	}

	// выводим сообщение в консоль
	fmt.Println("Recievtd message: ", msg.Message)

	sp, err := psql.OpenConn()
	if err != nil {
		log.Println(err.Error())
	}

	sp.Get(&storageData)
	fmt.Printf("Id: %d; Name: %s;", storageData.Id, storageData.Name)

	http.Redirect(w, r, "/", 302)

	// Handle JSON requests

	//err := json.NewDecoder(r.Body).Decode(&msg)
	//if err != nil {
	//	http.Error(w, "Invalid request", http.StatusBadRequest)
	//	return
	//}
	//defer r.Body.Close()
	//
	//fmt.Println("Received message:", msg.Message)
	//w.WriteHeader(http.StatusOK)
	//w.Write([]byte("Message received"))
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	//http.ServeFile(w, r, "./internal/html/main.html")
	tp, err := template.ParseFiles("./internal/html/main.html")
	if err != nil {
		log.Println(err.Error())
	}
	err = tp.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}
