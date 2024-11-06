package web

import (
	"fmt"

	// "fmt"
	"html/template"
	"log"
	"net/http"
	// "strconv"
)

type Message struct {
	Message string `json:"message"`
}

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
	fmt.Println("SendHandler function called")
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", 405)
	}

	// читаем параметр message из формы
	msg.Message = r.FormValue("message")

	// выводим сообщение в консоль
	fmt.Println("Recievtd message: ", msg.Message)

	// вернем ответ клиенту
	// if not send back, we will stay here
	//fmt.Fprintf(w, "Message recieved: %s", msg.Message)

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
