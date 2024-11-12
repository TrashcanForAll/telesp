package web

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strconv"
	"telesp/pkg/models"
	"telesp/pkg/models/psql"

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
func formJsonStr(data models.PersonData) []byte {
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println("jsonFile: ", jsonData)
	return jsonData
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
		data := models.PersonData{}
		respData := []models.PersonData{}

		tp, err := template.ParseFiles("./internal/html/main.html")
		if err != nil {
			log.Println(err.Error())
		}

		//for i := 0; i < countOfParams; i++ {
		data.FirstName = r.FormValue("field" + strconv.Itoa(1))
		data.LastName = r.FormValue("field" + strconv.Itoa(2))
		data.MiddleName = r.FormValue("field" + strconv.Itoa(3))
		data.Street = r.FormValue("field" + strconv.Itoa(4))
		data.House = r.FormValue("field" + strconv.Itoa(5))
		data.Building = r.FormValue("field" + strconv.Itoa(6))
		data.Apartment = r.FormValue("field" + strconv.Itoa(7))
		data.PhoneNumber = r.FormValue("field" + strconv.Itoa(8))
		fmt.Println("Igorava", r)
		jsonData := formJsonStr(data)

		resp, err := http.Post("http://127.0.0.1:8080/send", "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			log.Println(err.Error())
		}
		//fmt.Println("aboba was added")
		defer resp.Body.Close()
		//fmt.Println("response Status:", resp.Body)
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Read response error:", err)
			return
		}
		fmt.Println("Response: ", string(body))

		// Unmarshal
		if err := json.Unmarshal(body, &respData); err != nil {
			fmt.Println("Unmarshal response error:", err)
		}
		//print new params
		//fmt.Println(data.FirstName)
		//fmt.Println(data.LastName)

		showDataArr := []string{}
		for _, v := range respData {
			showDataArr = append(showDataArr, fmt.Sprintf("%v %v %v %v %v %v %v %v", v.FirstName, v.LastName, v.MiddleName, v.Street, v.House, v.Building, v.Apartment, v.PhoneNumber))
		}
		err = tp.Execute(w, showDataArr)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
		}
	}

}
