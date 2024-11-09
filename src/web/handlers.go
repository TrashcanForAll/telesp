package web

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strconv"
	"telesp/pkg/models"
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
	// test
	//p1 := r.FormValue("field1")
	//p2 := r.FormValue("field2")
	//fmt.Println("Aboba: ", p1, p2)

	var data models.TestPerson
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Prints
	fmt.Println(data.FirstName)
	fmt.Println(data.LastName)

	data.FirstName = data.FirstName + "aboba"
	data.LastName = data.LastName + "aboba"

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Println(err)
	}
	// выводим сообщение в консоль
	//fmt.Println("Recievtd message: ", msg.Message[0], msg.Message[1])

	//sp, err := psql.OpenConn()
	//if err != nil {
	//	log.Println(err.Error())
	//}

	//sp.Get(&storageData)
	//fmt.Printf("Id: %d; Name: %s;", storageData.Id, storageData.Name)

	//tp, err := template.ParseFiles("./internal/html/main.html")
	//if err != nil {
	//	log.Println(err.Error())
	//}
	//err = tp.Execute(w, msg.Message)
	//if err != nil {
	//	log.Println(err.Error())
	//	http.Error(w, "Internal Server Error", 500)
	//}
	//
	//http.Redirect(w, r, "/", 302)
}
func formJsonStr(data models.TestPerson) []byte {
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
		data := models.TestPerson{}
		respData := models.TestPerson{}

		tp, err := template.ParseFiles("./internal/html/main.html")
		if err != nil {
			log.Println(err.Error())
		}

		//for i := 0; i < countOfParams; i++ {
		data.FirstName = r.FormValue("field" + strconv.Itoa(1))
		data.LastName = r.FormValue("field" + strconv.Itoa(2))

		jsonData := formJsonStr(data)

		resp, err := http.Post("http://127.0.0.1:8080/send", "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			log.Println(err.Error())
		}
		fmt.Println("aboba was added")
		defer resp.Body.Close()
		fmt.Println("response Status:", resp.Body)
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
		fmt.Println(respData.FirstName)
		fmt.Println(respData.LastName)

		err = tp.Execute(w, []string{respData.FirstName, respData.LastName})
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
		}
	}

}
