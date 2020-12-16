package applications

import (
	"errors"
	"testing"

	"restapi/src/applications/mocks"
	"restapi/src/entities"
)

func TestCreateBookUsecaseWhenSuccessfully(t *testing.T) {
	sut := NewCreateBookUsecase(mocks.NewInMemoryRepository(mocks.ReturnCreate{}, mocks.ReturnFindByTitle{}, mocks.ReturnFindByID{}))

	result := sut.Create(&entities.BookDTO{
		Title:             "title",
		Author:            "author",
		PublishingCompany: "company",
		Edition:           1,
	})

	if result.StatusCode != 201 {
		t.Errorf("Should return status code 201 but received %d", result.StatusCode)
	}
}

func TestCreateBookUsecaseConflictErr(t *testing.T) {

	sut := NewCreateBookUsecase(mocks.NewInMemoryRepository(mocks.ReturnCreate{}, mocks.ReturnFindByTitle{Entity: &entities.BookEntity{}}, mocks.ReturnFindByID{}))

	result := sut.Create(&entities.BookDTO{
		Title:             "title",
		Author:            "author",
		PublishingCompany: "company",
		Edition:           1,
	})

	if result.StatusCode != 409 {
		t.Errorf("Should return status code 409 but received %d", result.StatusCode)
	}
}

func TestCreateBookUsecaseErrorOnFindByTitle(t *testing.T) {

	sut := NewCreateBookUsecase(mocks.NewInMemoryRepository(mocks.ReturnCreate{}, mocks.ReturnFindByTitle{Err: errors.New("")}, mocks.ReturnFindByID{}))

	result := sut.Create(&entities.BookDTO{
		Title:             "title",
		Author:            "author",
		PublishingCompany: "company",
		Edition:           1,
	})

	if result.StatusCode != 500 {
		t.Errorf("Should return status code 500 but received %d", result.StatusCode)
	}
}

func TestCreateBookUsecaseErrorOnCreate(t *testing.T) {

	sut := NewCreateBookUsecase(mocks.NewInMemoryRepository(mocks.ReturnCreate{Err: errors.New("")}, mocks.ReturnFindByTitle{}, mocks.ReturnFindByID{}))

	result := sut.Create(&entities.BookDTO{
		Title:             "title",
		Author:            "author",
		PublishingCompany: "company",
		Edition:           1,
	})

	if result.StatusCode != 500 {
		t.Errorf("Should return status code 500 but received %d", result.StatusCode)
	}
}
