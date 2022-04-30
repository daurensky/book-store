package book

import (
	"book-store/headers"
	"book-store/models"
	"book-store/state"
	"net/http"
	"strconv"
)

func Destroy(w http.ResponseWriter, r *http.Request) {
	headers.AllowOrigin(w)

	id := r.PostFormValue("Id")
	id_int, err := strconv.Atoi(id)

	if err != nil {
		panic(err)
	}

	books := []models.Book{}

	for _, book := range state.Books {
		if book.Id != id_int {
			books = append(books, book)
		}
	}

	state.Books = books
}
