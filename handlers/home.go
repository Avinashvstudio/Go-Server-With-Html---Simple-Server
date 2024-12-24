package handlers
import (
	"html/template"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request){
	tmpl, err := template.ParseFiles("templates/home.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	data := map[string]string{
		"title": "home page",
		"message": "welcome to the basic web project 2!",

	}
	tmpl.Execute(w , data)
}