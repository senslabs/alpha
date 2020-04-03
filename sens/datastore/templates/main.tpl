package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/senslabs/alpha/sens/datastore/ext"
	"github.com/senslabs/alpha/sens/datastore/generated/api"
	"github.com/senslabs/alpha/sens/logger"
)

func main() {
	r := mux.NewRouter()
	logger.InitLogger("sens.datastore")

	{{range .Models}}api.{{.}}Main(r)
	{{end}}
	ext.ExtMain(r)

	http.Handle("/", r)
	http.ListenAndServe(":9804", r)
}
