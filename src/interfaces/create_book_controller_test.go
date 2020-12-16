package interfaces

import (
	"net/http/httptest"
	"restapi/src/applications"
	"restapi/src/entities"
	"restapi/src/shared"
	"strings"
	"testing"
)

type createBookUsecaseMock struct {
	createdReturn *shared.HTTPResponse
}

func (mock *createBookUsecaseMock) Create(bookDTO *entities.BookDTO) *shared.HTTPResponse {
	return mock.createdReturn
}
func NewCreateBookUsecaseMock(createdReturn *shared.HTTPResponse) applications.ICreateBookUsecase {
	return &createBookUsecaseMock{createdReturn: createdReturn}
}

func TestCreateBookControllerSuccessfullyWithoutBody(t *testing.T) {

	req := httptest.NewRequest("POST", "http://example.com/foo",
		strings.NewReader(`{
			"title": "title",
			"author": "author",
			"publishingCompany": "publishingCompany",
			"edition": 1
	}`))
	res := httptest.NewRecorder()

	controller := CreateBookController{Usecase: NewCreateBookUsecaseMock(shared.HTTPCreated())}
	controller.Handler(res, req)

	result := res.Result()

	if result.StatusCode != 201 {
		t.Errorf("Should return status code 201 but received %d", result.StatusCode)
	}
}

func TestCreateBookControllerSuccessfullyWithBody(t *testing.T) {
	req := httptest.NewRequest("POST", "http://example.com/foo",
		strings.NewReader(`{
			"title": "title",
			"author": "author",
			"publishingCompany": "publishingCompany",
			"edition": 1
	}`))
	res := httptest.NewRecorder()

	controller := CreateBookController{Usecase: NewCreateBookUsecaseMock(shared.HTTPSuccess(""))}
	controller.Handler(res, req)

	result := res.Result()

	if result.StatusCode != 200 {
		t.Errorf("Should return status code 200 but received %d", result.StatusCode)
	}
}

func TestCreateBookControllerErrNoBody(t *testing.T) {
	req := httptest.NewRequest("POST", "http://example.com/foo", nil)
	res := httptest.NewRecorder()

	controller := CreateBookController{Usecase: NewCreateBookUsecaseMock(shared.HTTPCreated())}
	controller.Handler(res, req)

	result := res.Result()

	if result.StatusCode != 400 {
		t.Errorf("Should return status code 400 but received %d", result.StatusCode)
	}
}
func TestCreateBookControllerErrBodyWrongFormat(t *testing.T) {
	req := httptest.NewRequest("POST", "http://example.com/foo",
		strings.NewReader(`{
			"title": "title"
	}`))
	res := httptest.NewRecorder()

	controller := CreateBookController{Usecase: NewCreateBookUsecaseMock(shared.HTTPCreated())}
	controller.Handler(res, req)

	result := res.Result()

	if result.StatusCode != 400 {
		t.Errorf("Should return status code 400 but received %d", result.StatusCode)
	}
}
