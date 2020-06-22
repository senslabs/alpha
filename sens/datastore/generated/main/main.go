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
	api.BaselineMain(r)
	api.SessionEventMain(r)
	api.SessionRecordMain(r)
	api.SessionPropertieMain(r)
	api.AuthViewMain(r)
	api.OrgViewMain(r)
	api.OpViewMain(r)
	api.UserViewMain(r)
	api.OrgSessionRecordViewMain(r)
	api.DevicePropertieMain(r)
	api.OrgSessionEventViewMain(r)
	api.OrgSessionEventDetailViewMain(r)
	api.OrgSessionDetailViewMain(r)
	api.OrgPropertieMain(r)
	api.OpSettingMain(r)
	api.OpPropertieMain(r)
	api.OrgAlertViewMain(r)
	api.OrgLatestAlertViewMain(r)
	api.OrgSettingMain(r)
	api.UserSettingViewMain(r)
	api.BaselineViewMain(r)
	api.OrgSessionViewMain(r)
	api.UserSessionCountViewMain(r)
	api.SessionDurationViewMain(r)
	api.AlertRuleMain(r)
	api.AlertEscalationMain(r)
	api.ReportMain(r)
	api.ReportViewMain(r)
	api.LongestSleepTrendViewMain(r)
	api.UserDatedSessionViewMain(r)
	api.OrgSessionInfoViewMain(r)
	api.OrgMeditationViewMain(r)
	api.OrgSleepViewMain(r)
	api.SessionViewMain(r)
	api.UserBaselineViewMain(r)
	api.ResourceMain(r)
	api.SleepSummaryViewMain(r)
	api.MeditationSummaryViewMain(r)
	api.LatestSleepSummaryViewMain(r)
	api.LatestMeditationSummaryViewMain(r)
	api.DeviceViewMain(r)
	api.OrgActivityViewMain(r)
	api.OrgActivitySummaryViewMain(r)
	api.OrgQuarterUsageViewMain(r)
	api.RecorderMain(r)
	api.CollectorMain(r)
	api.CollectorPropertieMain(r)
	api.RecorderViewMain(r)

	ext.ExtMain(r)

	http.Handle("/", r)
	r.Use(loggingMiddleware)
	http.ListenAndServe(":9804", r)
}
