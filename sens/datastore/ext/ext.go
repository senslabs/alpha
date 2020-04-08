package ext

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/senslabs/alpha/sens/datastore"
	"github.com/senslabs/alpha/sens/httpclient"
	"github.com/senslabs/alpha/sens/logger"
	"github.com/senslabs/sqlx"
)

func ExtMain(r *mux.Router) {
	s := r.PathPrefix("/api/ext").Subrouter()
	s.HandleFunc("/activities/get", GetOrgActivites).Queries("days", "{days:[0-9]+}")
}

type Activity struct {
	Count        int    `db:"count"`
	ActivityType string `db:"activity_type"`
}

func GetOrgActivites(w http.ResponseWriter, r *http.Request) {
	in := r.URL.Query().Get("in")
	userIds := strings.Split(in, "^")
	if len(userIds) < 2 {
		logger.Error("Too less number of arguments")
		httpclient.WriteError(w, http.StatusInternalServerError, errors.New("Too less number of arguments"))
	} else {
		days := r.URL.Query().Get("days")
		values := map[string]interface{}{
			"user_ids": userIds[1:],
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
			if err := db.Select(&dest, query, args...); err != nil {
				logger.Error(err)
				httpclient.WriteError(w, http.StatusInternalServerError, err)
			} else if err := json.NewEncoder(w).Encode(dest); err != nil {
				logger.Error(err)
				httpclient.WriteError(w, http.StatusInternalServerError, err)
			}
		}
	}
}
