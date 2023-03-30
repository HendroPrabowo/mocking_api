package response

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func Ok(w http.ResponseWriter, resp interface{}) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	writeResponse(w, resp)
}

func Error(w http.ResponseWriter, status int, resp interface{}) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(status)
	writeResponse(w, resp)
}

func writeResponse(w http.ResponseWriter, resp interface{}) {
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Error(err)
	}
	w.Write(jsonResp)
}
