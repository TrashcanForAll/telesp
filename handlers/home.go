package handlers

import (
	"os"
	"net/http"
)
var hands map[string]func(w http.ResponseWriter, r *http.Request)

func HandlersLunch(path string) {
	
}

