package handlers

import "github.com/bond00729/info344-in-class/zipsvr/models"
import "net/http"
import "strings"
import "encoding/json"

type CityHandler struct {
	PathPrefix string
	Index      models.ZipIndex
}

// because we use ServeHTTP this becomes an interface usable in mux.Handle(path, handler)
// mux.HandleFunc(path, func) is for basic functions

// the ch allows for . notation similiar to java, allows us to access the stucts fields
func (ch *CityHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// URL: /zips/city-name
	cityName := r.URL.Path[len(ch.PathPrefix):]
	cityName = strings.ToLower(cityName)
	if len(cityName) == 0 {
		// http.StatusBadRequest = 400
		http.Error(w, "please provide a city name", http.StatusBadRequest)
	}

	w.Header().Add(headerContentType, contentTypeJSON)
	w.Header().Add(accessControlAllowOrigin, "*")
	zips := ch.Index[cityName]
	json.NewEncoder(w).Encode(zips)
}
