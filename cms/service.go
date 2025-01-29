package cms

import (
	"fmt"
	"net/http"
	"strings"

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
	entity.fromDto(dto)

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
	if dto.Response == nil {
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
	entity.fromDto(dto)

	if err := svc.repository.update(entity); err != nil {
		return wraped_error.WrapError(err, http.StatusInternalServerError)
	}

	return nil
}

func (svc service) proceedDeleteMock(id int) *wraped_error.Error {
	if err := svc.repository.delete(id); err != nil {
		return wraped_error.WrapError(err, http.StatusInternalServerError)
	}
	return nil
}

func (svc service) proceedGetMock(queryParam MockQueryDto) (mockResponseDto MockResponseDto, errWrap *wraped_error.Error) {
	mocksEntity, err := svc.repository.get(queryParam)
	if err != nil {
		return mockResponseDto, wraped_error.WrapError(err, http.StatusInternalServerError)
	}
	totalElement, err := svc.repository.count(queryParam)
	if err != nil {
		return mockResponseDto, wraped_error.WrapError(err, http.StatusInternalServerError)
	}
	mocksDto := []MockDto{}
	for _, mock := range mocksEntity {
		var dtoMock MockDto
		copier.Copy(&dtoMock, &mock)
		mocksDto = append(mocksDto, dtoMock)
	}

	mockResponseDto.Mock = mocksDto
	mockResponseDto.Page = queryParam.Page
	mockResponseDto.TotalElement = totalElement
	return mockResponseDto, nil
}

func (svc service) proceedHandleMock(method string, path string) (mockDto MockDto, errWrap *wraped_error.Error) {
	mockEntity, err := svc.repository.getOneByMethodAndPath(method, path)
	if err != nil {
		return MockDto{}, wraped_error.WrapError(err, http.StatusInternalServerError)
	}
	mockDto.fromEntity(mockEntity)
	return mockDto, nil
}
