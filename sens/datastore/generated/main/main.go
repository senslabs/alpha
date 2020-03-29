package main

import (
	"net/http"

	"github.com/gorilla/mux"
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
	api.EndpointMain(r)
	api.DeviceMain(r)
	api.OrgOpMain(r)
	api.OrgUserMain(r)
	api.OpUserCategorieMain(r)
	api.OpUserMain(r)
	api.OrgEndpointCategorieMain(r)
	api.OrgEndpointMain(r)
	api.OpEndpointCategorieMain(r)
	api.OpEndpointMain(r)
	api.UserEndpointCategorieMain(r)
	api.UserEndpointMain(r)
	api.SessionMain(r)
	api.SessionEventMain(r)
	api.SessionRecordMain(r)
	api.SessionPropertieMain(r)
	api.OrgDetailViewMain(r)
	api.OpDetailViewMain(r)
	api.UserDetailViewMain(r)
	api.DeviceViewMain(r)
	
	http.Handle("/", r)
	http.ListenAndServe(":9804", r)
}
