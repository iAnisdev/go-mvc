package controllers

import (
	"io"
	"net/http"
)

func LoginRoute(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "This is login route!\n")
}
