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

	api.DeviceActivitieMain(r)
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
	api.OrgEndpointAccessGroupMain(r)
	api.OrgEndpointMain(r)
	api.EndpointMain(r)
	api.OpEndpointAccessGroupMain(r)
	api.OpEndpointMain(r)
	api.UserEndpointAccessGroupMain(r)
	api.UserEndpointMain(r)
	api.DeviceMain(r)
	api.AlertMain(r)
	api.SessionMain(r)
	api.SessionSettingMain(r)
	api.VitalBaselineMain(r)
	api.SessionEventMain(r)
	api.SessionRecordMain(r)
	api.SessionPropertieMain(r)
	api.AuthDetailViewMain(r)
	api.OrgDetailViewMain(r)
	api.OpDetailViewMain(r)
	api.UserDetailViewMain(r)
	api.DeviceViewMain(r)
	api.UserSessionViewMain(r)
	api.UserAlertViewMain(r)
	api.SleepViewMain(r)
	api.MeditationViewMain(r)
	api.SleepSummarieMain(r)
	api.MeditationSummarieMain(r)
	api.UserSummaryViewMain(r)
	api.UserSleepViewMain(r)
	api.UserMeditationViewMain(r)
	
	ext.ExtMain(r)

	http.Handle("/", r)
	http.ListenAndServe(":9804", r)
}
