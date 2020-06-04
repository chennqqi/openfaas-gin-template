package function

import (
	"net/http"
)

var (
	app http.Handler
)

func Handle(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
