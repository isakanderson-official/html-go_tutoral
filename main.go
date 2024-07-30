package main

import (
	"log"
	"net/http"
	"text/template"
)

type Todo struct {
	Id int
	Message string
}

func main() {

	data := map[string][]Todo{
		"Todos": {
			Todo{Id:1, Message: "Buy Milk"},
		},
	}
	
	 todosHandler := func(w http.ResponseWriter, r *http.Request){
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.Execute(w, data)
	 }


	http.HandleFunc("/", todosHandler)

	log.Fatal(http.ListenAndServe(":8000", nil))
}