package web

import (
	"encoding/json"
	"fmt"
	"log"
	"telesp/pkg/models"
	"telesp/pkg/models/psql"
	sh "telesp/src/web/subhendlers"

	// "fmt"
	"html/template"
	"log/slog"
	"net/http"
	// "strconv"
)

const countOfParams = 2

type Message struct {
	Message [countOfParams]string `json:"message"`
}

var storageData = models.TestPerson{}
var logger = slog.Logger{}
var msg Message

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
	log.Println("SendHandler function called")
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", 405)
	}

	var data models.PersonData
	var respData = []models.PersonData{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	db, err := psql.OpenConn()
	if err != nil {
		log.Println(err)
	}
	defer db.DB.Close()

	respData = db.Get(&data)
	for _, v := range respData {
		fmt.Println(v)
	}
	// Prints
	//fmt.Println(respData[0].FirstName)
	//fmt.Println(data.LastName)
	//
	//data.FirstName = data.FirstName + "aboba"
	//data.LastName = data.LastName + "aboba"

	//for i, _ := range respData {
	err = json.NewEncoder(w).Encode(respData)
	if err != nil {
		log.Println(err)
	}
	//}

}

func AddHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", 405)
	}
	var data models.PersonData
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	db, err := psql.OpenConn()
	if err != nil {
		log.Println(err)
	}
	defer db.DB.Close()

	err = db.Insert(&data)
	if err != nil {
		log.Println(err)
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	//http.ServeFile(w, r, "./internal/html/main.html")

	if r.Method == http.MethodGet {
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
	if r.Method == http.MethodPost {
		if r.FormValue("search") == "Select" {
			tp, err := template.ParseFiles("./internal/html/main.html")
			if err != nil {
				log.Println(err.Error())
			}

			showDataArr := sh.SelectHandler(r)

			err = tp.Execute(w, showDataArr)
			if err != nil {
				log.Println(err.Error())
				http.Error(w, "Internal Server Error", 500)
			}

		}
		if r.FormValue("add") == "Add" {

			sh.AddHandler(r)

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
		if r.FormValue("update") == "Update" {
		}
		if r.FormValue("delete") == "Delete" {
		}
	}

}
