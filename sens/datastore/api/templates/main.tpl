package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	{{range .Models}}{{.}}Main(r)
	{{end}}
	http.Handle("/", r)
}
