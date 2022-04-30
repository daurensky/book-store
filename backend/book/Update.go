package book

import (
	"book-store/headers"
	"book-store/models"
	"book-store/state"
	"net/http"
	"strconv"
)

func Update(w http.ResponseWriter, r *http.Request) {
	headers.AllowOrigin(w)

	id := r.PostFormValue("Id")
	name := r.PostFormValue("Name")
	cost := r.PostFormValue("Cost")

	cost_float, err := strconv.ParseFloat(cost, 64)

	if err != nil {
		panic(err)
	}

	id_int, err := strconv.Atoi(id)

	if err != nil {
		panic(err)
	}

	books := []models.Book{}

	for _, book := range state.Books {
		if book.Id == id_int {
			books = append(books, models.Book{
				Id:   id_int,
				Name: name,
				Cost: cost_float,
			})
		} else {
			books = append(books, book)
		}
	}

	state.Books = books
}
