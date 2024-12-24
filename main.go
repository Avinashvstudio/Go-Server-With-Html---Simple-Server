package main
import (
	"fmt"
	// "html/template"
	"net/http"
	"basic-web-2/handlers"
)

func main(){
	http.HandleFunc("/",handlers.HomeHandler)
	http.HandleFunc("/about",  handlers.AboutHandler)


	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/static/",http.StripPrefix("/static",fs))

	fmt.Println("Server running on http://localhost:8080")
	err :=http.ListenAndServe(":8080" , nil)
	if err!= nil {
		fmt.Println("Error starting server:" , err)
	}
}