package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/senslabs/alpha/sens/datastore/generated/api"
	"github.com/senslabs/alpha/sens/logger"
)

func main() {
	r := mux.NewRouter()
	logger.InitLogger("")

	{{range .Models}}api.{{.}}Main(r)
	{{end}}
	http.Handle("/", r)
	http.ListenAndServe(":9804", r)
}
