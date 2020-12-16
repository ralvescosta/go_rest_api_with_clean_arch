package interfaces

import (
	"net/http/httptest"
	"restapi/src/applications"
	"restapi/src/entities"
	"restapi/src/shared"
	"testing"
)

type getOneBookUsecaseMock struct {
	getOneReturn *shared.HTTPResponse
}

func (mock *getOneBookUsecaseMock) GetOne(id uint64) *shared.HTTPResponse {
	return mock.getOneReturn
}
func NewGetOnBookUsecaseMock(getOneReturn *shared.HTTPResponse) applications.IGetOneBookUsecase {
	return &getOneBookUsecaseMock{getOneReturn: getOneReturn}
}

func TestGetOneBookControllerSuccessfully(t *testing.T) {

	req := httptest.NewRequest("GET", "http://example.com/1", nil)
	res := httptest.NewRecorder()

	controller := GetOneBookController{Usecase: NewGetOnBookUsecaseMock(shared.HTTPSuccess(entities.BookEntity{}))}
	controller.Handler(res, req)

	result := res.Result()

	if result.StatusCode != 200 {
		t.Errorf("Should return status code 200 but received %d", result.StatusCode)
	}

}
