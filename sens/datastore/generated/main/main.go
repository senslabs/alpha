package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/senslabs/alpha/sens/datastore/generated/api"
	"github.com/senslabs/alpha/sens/logger"
)

func main() {
	r := mux.NewRouter()
	logger.InitLogger("datastore")

	api.AuthMain(r)
	api.OrgMain(r)
	api.OpMain(r)
	api.UserMain(r)
	api.EndpointMain(r)
	api.DeviceMain(r)
	api.OrgAuthMain(r)
	api.OpAuthMain(r)
	api.UserAuthMain(r)
	api.OrgOpMain(r)
	api.OrgUserMain(r)
	api.OpUserMain(r)
	api.OrgEndpointMain(r)
	api.OpEndpointMain(r)
	api.UserEndpointMain(r)
	api.OrgDetailMain(r)
	api.OpDetailMain(r)
	api.UserDetailMain(r)
	api.SessionMain(r)
	api.SessionEventMain(r)
	api.SessionRecordMain(r)
	api.SessionPropertieMain(r)
	
	http.Handle("/", r)
	http.ListenAndServe(":9804", r)
}
