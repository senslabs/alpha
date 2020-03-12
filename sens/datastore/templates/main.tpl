package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	logger.InitFileLogger("datastore")

	{{range .Models}}{{.}}Main(r)
	{{end}}
	http.Handle("/", r)
	http.ListenAndServe(":9804", r)
}
