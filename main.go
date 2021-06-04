package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("Template/index.html")
	if err != nil {
		fmt.Fprintf(w, "Error al encontrar la pagina")
	} else {
		template.Execute(w, nil)
	}
}

func main() {

	http.HandleFunc("/", index)

	fmt.Println("App en ejecutandose")
	http.ListenAndServe(":3000", nil)
}
