package interfaces

import (
	"net/http/httptest"
	"restapi/src/applications"
	"restapi/src/entities"
	"restapi/src/shared"
	"strings"
	"testing"
)

type createBookUsecaseMock struct{}

func (*createBookUsecaseMock) Create(bookDTO *entities.BookDTO) *shared.HTTPResponse {
	return shared.HTTPCreated()
}
func NewCreateBookUsecaseMock() applications.ICreateBookUsecase {
	return &createBookUsecaseMock{}
}

func TestCreateBookControllerSuccessfully(t *testing.T) {

	req := httptest.NewRequest("POST", "http://example.com/foo",
		strings.NewReader(`{
			"title": "title",
			"author": "author",
			"publishingCompany": "publishingCompany",
			"edition": 1
	}`))
	res := httptest.NewRecorder()

	controller := CreateBookController{Usecase: NewCreateBookUsecaseMock()}
	controller.Handler(res, req)

	result := res.Result()

	if result.StatusCode != 201 {
		t.Errorf("Should return status code 201 but received %d", result.StatusCode)
	}
}
