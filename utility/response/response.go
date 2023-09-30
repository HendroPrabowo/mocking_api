package response

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"

	"mocking_api/utility/wraped_error"
)

func Ok(w http.ResponseWriter, resp interface{}) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	writeResponse(w, resp)
}

func OkWithMessage(w http.ResponseWriter, resp interface{}) {
	mapResp := map[string]interface{}{"message": resp}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	writeResponse(w, mapResp)
}

func Custom(w http.ResponseWriter, status int, resp interface{}) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(status)
	writeResponse(w, resp)
}

func ErrorWrapped(w http.ResponseWriter, err *wraped_error.Error) {
	mapResp := map[string]interface{}{"message": err.Err.Error()}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(err.StatusCode)
	writeResponse(w, mapResp)
}

func writeResponse(w http.ResponseWriter, resp interface{}) {
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Error(err)
	}
	w.Write(jsonResp)
}
