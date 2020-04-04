package main

import (
	"log"
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
	api.DeviceActivitieMain(r)
	api.AlertMain(r)
	api.SessionMain(r)
	api.SessionEventMain(r)
	api.SessionRecordMain(r)
	api.OrgDetailViewMain(r)
	api.OpDetailViewMain(r)
	api.UserDetailViewMain(r)
	api.DeviceViewMain(r)
	api.UserAlertViewMain(r)
	api.SleepViewMain(r)
	api.MeditationViewMain(r)
	api.UserSessionViewMain(r)
	api.SessionPropertieMain(r)
	api.SleepSummarieMain(r)
	api.MeditationSummarieMain(r)
	api.UserSummaryViewMain(r)
	api.UserSleepViewMain(r)
	api.UserMeditationViewMain(r)

	ext.ExtMain(r)

	http.Handle("/", r)
	log.Println(http.ListenAndServe(":9804", r))
}
