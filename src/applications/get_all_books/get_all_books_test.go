// package applications

// import (
// 	"errors"
// 	"restapi/src/applications/protocols"
// 	"restapi/src/entities"
// 	"testing"
// )

// type inMemoryRepository struct {
// 	returnCreate      returnCreate
// 	returnFindByTitle returnFindByTitle
// 	returnFindByID    returnFindByID
// }

// func TestCreateBookUsecaseWhenSuccessfully(t *testing.T) {
// 	sut := NewCreateBookUsecase(NewInMemoryRepository(returnCreate{}, returnFindByTitle{}, returnFindByID{}))

// 	result := sut.Create(&entities.BookDTO{
// 		Title:             "title",
// 		Author:            "author",
// 		PublishingCompany: "company",
// 		Edition:           1,
// 	})

// 	if result.StatusCode != 201 {
// 		t.Errorf("Should return status code 201 but received %d", result.StatusCode)
// 	}
// }

// func TestCreateBookUsecaseConflictErr(t *testing.T) {

// 	sut := NewCreateBookUsecase(NewInMemoryRepository(returnCreate{}, returnFindByTitle{entity: &entities.BookEntity{}}, returnFindByID{}))

// 	result := sut.Create(&entities.BookDTO{
// 		Title:             "title",
// 		Author:            "author",
// 		PublishingCompany: "company",
// 		Edition:           1,
// 	})

// 	if result.StatusCode != 409 {
// 		t.Errorf("Should return status code 409 but received %d", result.StatusCode)
// 	}
// }

// func TestCreateBookUsecaseErrorOnFindByTitle(t *testing.T) {

// 	sut := NewCreateBookUsecase(NewInMemoryRepository(returnCreate{}, returnFindByTitle{err: errors.New("")}, returnFindByID{}))

// 	result := sut.Create(&entities.BookDTO{
// 		Title:             "title",
// 		Author:            "author",
// 		PublishingCompany: "company",
// 		Edition:           1,
// 	})

// 	if result.StatusCode != 500 {
// 		t.Errorf("Should return status code 500 but received %d", result.StatusCode)
// 	}
// }

// func TestCreateBookUsecaseErrorOnCreate(t *testing.T) {

// 	sut := NewCreateBookUsecase(NewInMemoryRepository(returnCreate{err: errors.New("")}, returnFindByTitle{}, returnFindByID{}))

// 	result := sut.Create(&entities.BookDTO{
// 		Title:             "title",
// 		Author:            "author",
// 		PublishingCompany: "company",
// 		Edition:           1,
// 	})

// 	if result.StatusCode != 500 {
// 		t.Errorf("Should return status code 500 but received %d", result.StatusCode)
// 	}
// }
