package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	http.HandleFunc("/hello", sayHello)

	log.Println("App started...")
	http.ListenAndServe(":"+port, nil)
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello!"))
}
