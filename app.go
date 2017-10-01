package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	http.HandleFunc("/hello", sayHello)

	http.HandleFunc("/", index)
	http.HandleFunc("/product", product)

	log.Println("App started...")
	http.ListenAndServe(":"+port, nil)
}

// sayHello will directly write to response writer
func sayHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello!"))
}

// index will parse template file and write it to response writer
func index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, nil)
}

type ProductData struct {
	ID       int
	Nama     string
	Harga    string
	Berat    float64
	BeratFmt string
}

// product will parse template file and write products data to response writer
func product(w http.ResponseWriter, r *http.Request) {
	berat := 1200000.512312310293090193109310
	productList := []ProductData{
		ProductData{
			ID:       1,
			Nama:     "Laptop",
			Harga:    "10.000.000",
			Berat:    berat,
			BeratFmt: strconv.FormatFloat(berat, 'f', -1, 64),
		},
		ProductData{
			ID:       2,
			Nama:     "Keyboard",
			Harga:    "10.000",
			Berat:    berat,
			BeratFmt: strconv.FormatFloat(berat, 'f', -1, 64),
		},
	}

	data := productList
	t, err := template.ParseFiles("templates/product.html")
	if err != nil {
		log.Println(err)
		return
	}
	t.Execute(w, data)
}
