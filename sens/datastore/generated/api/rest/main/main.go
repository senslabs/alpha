package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/senslabs/alpha/sens/logger"
)

func main() {
	r := mux.NewRouter()
	logger.InitLogger("datastore")

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
	SessionMain(r)
	SessionEventMain(r)
	SessionRecordMain(r)
	SessionPropertieMain(r)
	
	http.Handle("/", r)
	http.ListenAndServe(":9804", r)
}
