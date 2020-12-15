package interfaces

import (
	"fmt"
	"net/http/httptest"
	"restapi/src/applications"
	"restapi/src/entities"
	"restapi/src/shared"
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
	req := httptest.NewRequest("POST", "http://example.com/foo", nil)
	res := httptest.NewRecorder()

	controller := CreateBookController{Usecase: NewCreateBookUsecaseMock()}
	controller.Handler(res, req)

	result := res.Result()
	fmt.Println(result.StatusCode)
	if result.StatusCode != 201 {
		t.Errorf("")
	}
}
