package cms

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-pg/pg/v10"
	log "github.com/sirupsen/logrus"

	"mocking_api/utility/response"
	"mocking_api/utility/wraped_error"
)

type controller struct {
	service service
}

func newController(service service) controller {
	return controller{
		service: service,
	}
}

func (c controller) GetMock(w http.ResponseWriter, r *http.Request) {
	queryParam, err := c.buildLoanParameter(r)
	if err != nil {
		response.ErrorWrapped(w, err)
		return
	}

	mockResponseDto, err := c.service.proceedGetMock(queryParam)
	if err != nil {
		response.ErrorWrapped(w, err)
		return
	}

	response.Ok(w, mockResponseDto)
}

func (c controller) buildLoanParameter(r *http.Request) (queryParam MockQueryDto, wrappedError *wraped_error.Error) {
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")
	name := r.URL.Query().Get("name")
	method := r.URL.Query().Get("method")
	path := r.URL.Query().Get("path")
	responseCode := r.URL.Query().Get("response_code")

	if page == "" {
		page = "0"
	}

	if limit == "" {
		limit = "10"
	}

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return queryParam, wraped_error.WrapError(fmt.Errorf("invalid value page"), http.StatusBadRequest)
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		return queryParam, wraped_error.WrapError(fmt.Errorf("invalid value limit"), http.StatusBadRequest)
	}

	if responseCode != "" {
		responseCodeInt, err := strconv.Atoi(responseCode)
		if err != nil {
			return MockQueryDto{}, wraped_error.WrapError(fmt.Errorf("invalid value response_code"), http.StatusBadRequest)
		}
		queryParam.ResponseCode = responseCodeInt
	}

	queryParam.Page = pageInt
	queryParam.Limit = limitInt
	queryParam.Name = name
	queryParam.Method = method
	queryParam.Path = path

	return queryParam, nil
}

func (c controller) AddMock(w http.ResponseWriter, r *http.Request) {
	var dto MockDto
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		errWrap := wraped_error.WrapError(err, http.StatusBadRequest)
		response.ErrorWrapped(w, errWrap)
		return
	}

	if err := c.service.proceedAddMock(dto); err != nil {
		response.ErrorWrapped(w, err)
		return
	}

	response.OkWithMessage(w, "ok")
}

func (c controller) UpdateMock(w http.ResponseWriter, r *http.Request) {
	var dto MockDto
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		errWrap := wraped_error.WrapError(err, http.StatusBadRequest)
		response.ErrorWrapped(w, errWrap)
		return
	}

	if err := c.service.proceedUpdateMock(dto); err != nil {
		response.ErrorWrapped(w, err)
		return
	}

	response.OkWithMessage(w, "ok")
}

func (c controller) DeleteMock(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		response.ErrorWrapped(w, wraped_error.WrapError(err, http.StatusBadRequest))
		return
	}

	if err := c.service.proceedDeleteMock(id); err != nil {
		response.ErrorWrapped(w, err)
		return
	}
	response.OkWithMessage(w, "success")
}

func (c controller) HandleMock(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	method := strings.ToUpper(r.Method)
	log.Infof("%s : %s", method, path)
	mock, err := c.service.proceedHandleMock(method, path)
	if err != nil {
		if err.Err == pg.ErrNoRows {
			response.ErrorWrapped(w, wraped_error.WrapError(fmt.Errorf("mock not found"), http.StatusNotFound))
			return
		}
		response.ErrorWrapped(w, err)
		return
	}

	response.Custom(w, mock.ResponseCode, mock.Response)
}
