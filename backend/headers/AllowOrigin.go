package headers

import "net/http"

func AllowOrigin(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
}
