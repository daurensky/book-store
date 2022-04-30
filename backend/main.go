package main

import (
	"book-store/book"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/book", book.Index)
	http.HandleFunc("/book/store", book.Store)
	http.HandleFunc("/book/update", book.Update)
	http.HandleFunc("/book/destroy", book.Destroy)

	err := http.ListenAndServe("localhost:7777", http.DefaultServeMux)

	if err != nil {
		log.Fatal(err)
	}
}
