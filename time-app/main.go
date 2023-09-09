package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
	_ "time/tzdata"
)

func main() {
	log.Println("Initialized app")
	http.HandleFunc("/", index)

	address := ":" + os.Getenv("PORT")
	log.Println("Starting server at address", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func index(w http.ResponseWriter, req *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, time.Now().Format("15:04:05"))
}
