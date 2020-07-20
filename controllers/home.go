package controllers

import (
	"net/http"
)

//Home endpoint handler
var Home GetOnlyHandlerFunc = func(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/index.html")
}
