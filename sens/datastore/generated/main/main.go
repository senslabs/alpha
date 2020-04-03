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

	api.AuthMain(r)
	api.OrgMain(r)
	api.OpMain(r)
	api.UserMain(r)
	api.OpUserAccessGroupMain(r)
	api.OpUserMain(r)
	api.EndpointMain(r)
	api.OrgEndpointAccessGroupMain(r)
	api.OrgEndpointMain(r)
	api.OpEndpointAccessGroupMain(r)
	api.OpEndpointMain(r)
	api.UserEndpointAccessGroupMain(r)
	api.UserEndpointMain(r)
	api.DeviceMain(r)
	api.SessionMain(r)
	api.SessionEventMain(r)
	api.SessionRecordMain(r)
	api.SessionPropertieMain(r)
	api.OrgDetailViewMain(r)
	api.OpDetailViewMain(r)
	api.UserDetailViewMain(r)
	api.DeviceViewMain(r)

	ext.ExtMain(r)

	http.Handle("/", r)
	http.ListenAndServe(":9804", r)
}
