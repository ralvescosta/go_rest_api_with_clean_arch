package interfaces

import (
	"net/http/httptest"
	"restapi/src/applications"
	"restapi/src/entities"
	"restapi/src/shared"
	"testing"
)

type getAllBooksUsecaseMock struct {
	getAllReturn *shared.HTTPResponse
}

func (mock *getAllBooksUsecaseMock) GetAll() *shared.HTTPResponse {
	return mock.getAllReturn
}
func NewGetAllBooksUsecaseMock(getAllReturn *shared.HTTPResponse) applications.IGetAllBooksUsecase {
	return &getAllBooksUsecaseMock{getAllReturn: getAllReturn}
}

func TestGetAllBooksControllerSuccessfully(t *testing.T) {

	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	res := httptest.NewRecorder()

	controller := GetAllBooksController{Usecase: NewGetAllBooksUsecaseMock(shared.HTTPSuccess(entities.BookEntity{}))}
	controller.Handler(res, req)

	result := res.Result()

	if result.StatusCode != 200 {
		t.Errorf("Should return status code 200 but received %d", result.StatusCode)
	}
}
