package interfaces

import (
	"net/http/httptest"
	"restapi/src/applications"
	"restapi/src/shared"
	"testing"
)

type healthUsecaseMock struct{}

func (healthUsecaseMock) Health() *shared.HTTPResponse {
	return shared.HTTPSuccess("")
}
func NewHealthUsecaseMock() applications.IHealthUsecase {
	return &healthUsecaseMock{}
}

func TestHealthControllerSuccessfully(t *testing.T) {
	req := httptest.NewRequest("POST", "http://example.com/foo", nil)
	res := httptest.NewRecorder()

	controller := HealthController{Usecase: NewHealthUsecaseMock()}
	controller.Handler(res, req)

	result := res.Result()

	if result.StatusCode != 200 {
		t.Errorf("Should return status code 200 but received %d", result.StatusCode)
	}
}
