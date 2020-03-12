package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/senslabs/alpha/sens/logger"
)

func main() {
	r := mux.NewRouter()

	AuthMain(r)
	OrgMain(r)
	OpMain(r)
	UserMain(r)
	EndpointMain(r)
	OrgAuthMain(r)
	OpAuthMain(r)
	UserAuthMain(r)
	OrgOpMain(r)
	OrgUserMain(r)
	OpUserMain(r)
	OrgEndpointMain(r)
	OpEndpointMain(r)
	UserEndpointMain(r)

	logger.InitFileLogger("datastore")
	http.Handle("/", r)
	http.ListenAndServe(":9804", r)
}
