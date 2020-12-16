package applications

import (
	"restapi/src/applications/mocks"
	"restapi/src/entities"
	"testing"
)

func TestGetOneBookUsecaseSuccessfully(t *testing.T) {
	sut := NewGetOneBookUsecase(mocks.NewInMemoryRepository(mocks.ReturnCreate{}, mocks.ReturnFindByTitle{}, mocks.ReturnFindByID{Entity: &entities.BookEntity{}}, mocks.ReturnFindAll{}))

	result := sut.GetOne(1)

	if result.StatusCode != 200 {
		t.Errorf("Should return status code 200 but received %d", result.StatusCode)
	}
}
