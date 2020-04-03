package ext

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/senslabs/alpha/sens/datastore"
	"github.com/senslabs/alpha/sens/httpclient"
	"github.com/senslabs/alpha/sens/logger"
)

func ExtMain(r *mux.Router) {
	s := r.PathPrefix("/api/ext").Subrouter()
	s.HandleFunc("/activities/get", GetOrgActivites).Queries("days", "{days:[0-9]+}")
}

type Activity struct {
	Count int    `db:"count"`
	Type  string `db:"type"`
}

func GetOrgActivites(w http.ResponseWriter, r *http.Request) {
	userIds := r.URL.Query()["UserId"]
	days := r.URL.Query().Get("days")
	values := map[string]interface{}{
		"user_ids": userIds,
		"when":     days,
	}
	if query, args, err := sqlx.Named(ACTIVITY_DASHBOARD_QUERY, values); err != nil {
		logger.Error(err)
		httpclient.WriteError(w, http.StatusInternalServerError, err)
	} else if query, args, err := sqlx.In(query, args...); err != nil {
		logger.Error(err)
		httpclient.WriteError(w, http.StatusInternalServerError, err)
	} else {
		db := datastore.GetConnection()
		query = db.Rebind(query)
		logger.Debug(query, args)
		var dest []Activity
		logger.Error(db.Select(&dest, query, args...))
		logger.Debug(dest)
	}
}
