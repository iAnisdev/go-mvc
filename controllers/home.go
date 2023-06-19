package controllers

import (
	"io"
	"net/http"
)

func HomeRoute(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "This is my home route!\n")
}
