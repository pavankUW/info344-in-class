package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"../models"
)

type CityHandler struct {
	PathPrefix string
	Index      models.ZipIndex
}

func (ch *CityHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// URL: /zips/city-name
	cityName := r.URL.Path[len(ch.PathPrefix):]
	cityName = strings.ToLower(cityName)
	if len(cityName) == 0 {
		http.Error(w, "please provide a city name", http.StatusBadRequest)
		return
	}

	w.Header().Add(headerContentType, contentTypeJSON)
	zips := ch.Index[cityName]
	json.NewEncoder(w).Encode(zips)
}
