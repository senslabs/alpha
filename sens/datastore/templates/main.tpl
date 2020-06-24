package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/senslabs/alpha/sens/datastore/ext"
	"github.com/senslabs/alpha/sens/datastore/generated/api"
	"github.com/senslabs/alpha/sens/logger"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Debug(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func main() {
	r := mux.NewRouter()
	logger.InitLogger("sens.datastore")

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Hello Health!")
	})

	{{range .Models}}api.{{.}}Main(r)
	{{end}}
	ext.ExtMain(r)

	http.Handle("/", r)
	r.Use(loggingMiddleware)
	http.ListenAndServe(":9804", r)
}
