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
	api.SurveyQuestionMain(r)
	api.SurveyAnswerMain(r)
	api.UserSettingMain(r)
	api.UserPropertieMain(r)
	api.ApiKeyMain(r)
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
	api.SessionSettingMain(r)
	api.VitalBaselineMain(r)
	api.SessionEventMain(r)
	api.SessionRecordMain(r)
	api.SessionPropertieMain(r)
	api.AuthViewMain(r)
	api.OrgViewMain(r)
	api.OpViewMain(r)
	api.UserViewMain(r)
	api.DeviceViewMain(r)
	api.OrgAlertViewMain(r)
	api.OrgSessionViewMain(r)
	api.OrgActivityViewMain(r)
	api.OrgActivitySummaryViewMain(r)
	api.OrgQuarterUsageViewMain(r)
	api.OrgSessionRecordViewMain(r)
	api.DevicePropertieMain(r)
	api.OrgSessionInfoViewMain(r)
	api.OrgSleepViewMain(r)
	api.OrgMeditationViewMain(r)
	api.OrgSessionDetailViewMain(r)
	
	ext.ExtMain(r)

	http.Handle("/", r)
	http.ListenAndServe(":9804", r)
}
