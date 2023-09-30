package cms

import (
	"fmt"
	"net/http"
	"strings"

	"mocking_api/utility/time_now"
	"mocking_api/utility/wraped_error"

	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/copier"
)

type service struct {
	repository repository
}

func newService(repository repository) service {
	return service{
		repository: repository,
	}
}

func (svc service) proceedAddMock(dto MockDto) *wraped_error.Error {
	if err := svc.validateMockDto(dto); err != nil {
		return wraped_error.WrapError(err, http.StatusBadRequest)
	}

	entity := Mock{}
	copier.Copy(&entity, &dto)

	if err := svc.repository.insert(entity); err != nil {
		return wraped_error.WrapError(err, http.StatusInternalServerError)
	}

	return nil
}

func (svc service) validateMockDto(dto MockDto) error {
	validate := validator.New()
	if err := validate.Struct(dto); err != nil {
		return err
	}
	if dto.Response == "" {
		return fmt.Errorf("response cannot empty string")
	}
	dto.Method = strings.ToUpper(dto.Method)
	b, ok := mehtodAllowed[dto.Method]
	if !ok || !b {
		return fmt.Errorf("invalid method")
	}
	return nil
}

func (svc service) proceedUpdateMock(dto MockDto) *wraped_error.Error {
	if err := svc.validateMockDto(dto); err != nil {
		return wraped_error.WrapError(err, http.StatusBadRequest)
	}

	entity := Mock{}
	copier.Copy(&entity, &dto)
	entity.UpdatedAt = time_now.Wib().Format("2006-01-02 15:04:05")

	if err := svc.repository.update(entity); err != nil {
		return wraped_error.WrapError(err, http.StatusInternalServerError)
	}

	return nil
}

func (svc service) proceedGetMock(queryParam MockQueryDto) (mocksDto []MockDto, errWrap *wraped_error.Error) {
	mocksEntity, err := svc.repository.get(queryParam)
	if err != nil {
		return mocksDto, wraped_error.WrapError(err, http.StatusInternalServerError)
	}
	mocksDto = []MockDto{}
	for _, mock := range mocksEntity {
		var dtoMock MockDto
		copier.Copy(&dtoMock, &mock)
		mocksDto = append(mocksDto, dtoMock)
	}

	return mocksDto, nil
}

func (svc service) proceedHandleMock(method string, path string) (mockDto MockDto, errWrap *wraped_error.Error) {
	mockEntity, err := svc.repository.getOneByMethodAndPath(method, path)
	if err != nil {
		return MockDto{}, wraped_error.WrapError(err, http.StatusInternalServerError)
	}
	copier.Copy(&mockDto, &mockEntity)
	return mockDto, nil
}
