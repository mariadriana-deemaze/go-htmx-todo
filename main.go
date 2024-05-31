package main

import (
	"html/template"
	"log"
	"net/http"
)

type Todo struct {
	Id      int
	Message string
}

func main() {
	data := map[string][]Todo{
		"Todos": {
			Todo{Id: 1, Message: "Finish the other projects"},
		},
	}

	getTodos := func(w http.ResponseWriter, r *http.Request) {
		templ := template.Must(template.ParseFiles("index.html"))
		templ.Execute(w, data)

	}

	addTodo := func(w http.ResponseWriter, r *http.Request) {
		message := r.PostFormValue("message")
		templ := template.Must(template.ParseFiles("index.html"))
		todo := Todo{Id: len(data["Todos"]) + 1, Message: message}
		data["Todos"] = append(data["Todos"], todo)
		templ.ExecuteTemplate(w, "todo-list-element", todo)
	}

	http.HandleFunc("/", getTodos)
	http.HandleFunc("/add-todo", addTodo)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
