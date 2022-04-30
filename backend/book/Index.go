package book

import (
	"book-store/headers"
	"book-store/state"
	"encoding/json"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	headers.AllowOrigin(w)

	json_resp, err := json.Marshal(state.Books)

	if err != nil {
		panic(err)
	}

	w.Write(json_resp)
}
