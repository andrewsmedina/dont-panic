package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("ops!")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
