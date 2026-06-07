package main

import (
	"fmt"
	"net/http"
	"os"
)

func manoFunkcija(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Kristupas Vi Lambada funkcija 2026?")
}

func main() {

	http.HandleFunc("/", manoFunkcija)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Serveris paleistas ant render")

	http.ListenAndServe(":"+port, nil)
}
