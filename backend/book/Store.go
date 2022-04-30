package book

import (
	"book-store/headers"
	"book-store/models"
	"book-store/state"
	"net/http"
	"strconv"
)

func Store(w http.ResponseWriter, r *http.Request) {
	headers.AllowOrigin(w)

	name := r.PostFormValue("Name")
	cost := r.PostFormValue("Cost")

	cost_float, err := strconv.ParseFloat(cost, 64)

	if err != nil {
		panic(err)
	}

	id := 1

	if state.Books != nil && len(state.Books) > 0 {
		id = state.Books[len(state.Books)-1].Id + 1
	}

	state.Books = append(state.Books, models.Book{
		Id:   id,
		Name: name,
		Cost: cost_float,
	})
}
