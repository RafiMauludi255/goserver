package main

import (
	"fmt"
	"net/http"
)

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Ini halaman about")
}

func contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Ini halaman contact")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Halaman utama")

	})

	http.HandleFunc("/about", about)

	http.HandleFunc("/contact", contact)

	fmt.Println("Server running at http://localhost:8181")
	http.ListenAndServe(":8181", nil)
}
