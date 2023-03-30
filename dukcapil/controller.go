package dukcapil

import (
	"net/http"

	"mocking_api/utility/response"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

type controller struct {
}

func newController() controller {
	return controller{}
}

func (c controller) DukcapilIdentityVerify(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	status := r.URL.Query().Get("status")
	log.WithField("id", id).WithField("status", status).Info("dukcapil identity verify")

	dukcapilResponse := generateDukcapilSuccessResponse()
	if status == ERROR {
		dukcapilResponse = generateDukcapilFailResponse()
	}

	response.Ok(w, dukcapilResponse)
}

func generateDukcapilFailResponse() DukcapilResponse {
	content := Content{
		ResponseCode: "14",
		ResponseDesc: "Data Ditemukan, Status Tidak Aktif silakan hubungi Dinas Dukcapil",
		Response:     "Data Ditemukan, Status Tidak Aktif silakan hubungi Dinas Dukcapil",
	}
	response := DukcapilResponse{
		Content:          []Content{content},
		LastPage:         false,
		NumberOfElements: 0,
		Sort:             nil,
		TotalElements:    0,
		FirstPage:        false,
		Number:           0,
		Size:             1,
		QuotaLimiter:     1523,
	}
	return response
}

func generateDukcapilSuccessResponse() DukcapilResponse {
	content := Content{
		PlaceBirth:  "Sesuai (100)",
		MotherName:  "Sesuai (100)",
		Name:        "Sesuai (100)",
		DateOfBirth: "Sesuai",
		Gender:      "Sesuai",
	}
	response := DukcapilResponse{
		Content:          []Content{content},
		LastPage:         false,
		NumberOfElements: 0,
		Sort:             nil,
		TotalElements:    0,
		FirstPage:        false,
		Number:           0,
		Size:             1,
		QuotaLimiter:     1523,
	}
	return response
}
