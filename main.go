package main

import (
	"fmt"
	"net/http"
)

func manoFunkcija(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Testukas mano serveris veikia?")
}

func main() {

	http.HandleFunc("/", manoFunkcija)
	fmt.Println("Paspauskite ant linko http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}
