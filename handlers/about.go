package handlers
import (
	"html/template"
	"net/http"
)

func AboutHandler (w http.ResponseWriter, r *http.Request){
	tmpl,err :=template.ParseFiles("templates/about.html")
	if err !=nil{
		http.Error(w , "Error loading template", http.StatusInternalServerError)
		return
	}
	data := map[string]string{
		"title":"about Page",
		"Message":"This is the About page of our project!",
	
	}

	tmpl.Execute(w ,data)
}
