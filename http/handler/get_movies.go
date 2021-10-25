package handler

import (
	"encoding/json"
	"net/http"
	"stockbit-backend/cfg"
	"stockbit-backend/models/requests"
	"stockbit-backend/repository/logging"

	"github.com/gocraft/dbr"
	"github.com/spf13/viper"
)

type GetMovies struct {
	DBConn *dbr.Connection
}

func (c GetMovies) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

	var queryParameters requests.GetMoviesFilterRequest
	if r.FormValue("q") != "" {
		searchWord := r.FormValue("q")
		queryParameters.SearchWord = &searchWord
	}

	if r.FormValue("page") != "" {
		page := r.FormValue("page")
		queryParameters.Page = &page
	}

	var omdbParameter string
	if queryParameters.SearchWord != nil {
		omdbParameter = "&s=" + *queryParameters.SearchWord
	}
	if queryParameters.Page != nil {
		omdbParameter = omdbParameter + "&page=" + *queryParameters.Page
	}

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
