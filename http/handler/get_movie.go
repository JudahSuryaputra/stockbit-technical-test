package handler

import (
	"encoding/json"
	"net/http"
	"stockbit-backend/cfg"
	"stockbit-backend/models/requests"
	"stockbit-backend/repository/logging"

	"github.com/gocraft/dbr"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

type GetMovie struct {
	DBConn *dbr.Connection
}

func (c GetMovie) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sess := c.DBConn.NewSession(nil)

	log := requests.WriteLogRequest{
		Method:    r.Method,
		URL:       r.URL.RawPath,
		Host:      r.Host,
		UserAgent: r.UserAgent(),
	}

	err := logging.WriteLog(sess, log)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	pathVariable := mux.Vars(r)
	movieID := pathVariable["id"]

	omdbParameter := "&i=" + movieID

	url := viper.GetString(cfg.OmdbURL) + "/?" + viper.GetString(cfg.OmdbKey) + omdbParameter

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	defer response.Body.Close()

	var data map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
	return
}
