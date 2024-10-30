package web

import (
	// "fmt"
	"html/template"
	"log"
	"net/http"
	// "strconv"
)

func Home(w http.ResponseWriter, r *http.Request) {
	// Check if cur url match "/" template
	if r.URL.Path != "/" {
		http.NotFound(w,r)
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

	// w.Write([]byte("Привет из Snippetbox"))
}
 
// func ShowSnippet(w http.ResponseWriter, r *http.Request) {
// 	id, err := strconv.Atoi(r.URL.Query().Get("id"))
// 	if err != nil || id < 1 {
// 		http.NotFound(w, r)
// 		return
// 	}
	
// 	fmt.Fprintf(w, "Отображение выбранной заметки c ID: %d...", id)
// }
 

// func CreateSnippet(w http.ResponseWriter, r *http.Request) {
//     if r.Method != http.MethodPost {
// 		w.Header().Set("Allow", http.MethodPost)
// 		http.Error(w, "Метод запрещен!", 405)
//         return
//     }
 
//     w.Write([]byte("Создание новой заметки..."))
// }
 