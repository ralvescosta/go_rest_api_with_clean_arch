package applications

import (
	"testing"

	"restapi/src/applications/mocks"
)

func TestGetAllBooksUsecaseWhenSuccessfully(t *testing.T) {
	sut := NewGetAllBooksUsecase(mocks.NewInMemoryRepository(mocks.ReturnCreate{}, mocks.ReturnFindByTitle{}, mocks.ReturnFindByID{}, mocks.ReturnFindAll{}))

	result := sut.GetAll()

	if result.StatusCode != 200 {
		t.Errorf("Should return status code 200 but received %d", result.StatusCode)
	}
}
