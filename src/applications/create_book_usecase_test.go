package applications

import (
	"errors"
	"restapi/src/applications/protocols"
	"restapi/src/entities"
	"testing"
)

type inMemoryRepository struct {
	returnCreate      returnCreate
	returnFindByTitle returnFindByTitle
	returnFindByID    returnFindByID
}
type returnCreate struct {
	entity *entities.BookEntity
	err    error
}
type returnFindByTitle struct {
	entity *entities.BookEntity
	err    error
}
type returnFindByID struct {
	entity *entities.BookEntity
	err    error
}

func (repo *inMemoryRepository) Create(entity *entities.BookEntity) (*entities.BookEntity, error) {
	return repo.returnCreate.entity, repo.returnCreate.err
}
func (repo *inMemoryRepository) FindByTitle(title string) (*entities.BookEntity, error) {
	return repo.returnFindByTitle.entity, repo.returnFindByTitle.err
}
func (repo *inMemoryRepository) FindByID(id uint64) (*entities.BookEntity, error) {
	return repo.returnFindByID.entity, repo.returnFindByID.err
}
func NewInMemoryRepository(returnCreate returnCreate, returnFindByTitle returnFindByTitle, returnFindByID returnFindByID) protocols.IBooksRepository {
	return &inMemoryRepository{returnCreate: returnCreate, returnFindByTitle: returnFindByTitle, returnFindByID: returnFindByID}
}

func TestCreateBookUsecaseWhenSuccessfully(t *testing.T) {
	sut := CreateBookUsecase{BooksRepository: NewInMemoryRepository(returnCreate{}, returnFindByTitle{}, returnFindByID{})}

	result := sut.Create(&entities.BookDTO{
		Title:             "title",
		Author:            "author",
		PublishingCompany: "company",
		Edition:           1,
	})

	if result.StatusCode != 201 {
		t.Errorf("Should return status code %d, when is spected to be 201", result.StatusCode)
	}
}

func TestCreateBookUsecaseConflictErr(t *testing.T) {

	sut := CreateBookUsecase{BooksRepository: NewInMemoryRepository(returnCreate{}, returnFindByTitle{entity: &entities.BookEntity{}}, returnFindByID{})}

	result := sut.Create(&entities.BookDTO{
		Title:             "title",
		Author:            "author",
		PublishingCompany: "company",
		Edition:           1,
	})

	if result.StatusCode != 409 {
		t.Errorf("Should return status code %d, when is spected to be 409", result.StatusCode)
	}
}

func TestCreateBookUsecaseErrorOnFindByTitle(t *testing.T) {

	sut := CreateBookUsecase{BooksRepository: NewInMemoryRepository(returnCreate{}, returnFindByTitle{err: errors.New("")}, returnFindByID{})}

	result := sut.Create(&entities.BookDTO{
		Title:             "title",
		Author:            "author",
		PublishingCompany: "company",
		Edition:           1,
	})

	if result.StatusCode != 500 {
		t.Errorf("Should return status code %d, when is spected to be 500", result.StatusCode)
	}
}

func TestCreateBookUsecaseErrorOnCreate(t *testing.T) {

	sut := CreateBookUsecase{BooksRepository: NewInMemoryRepository(returnCreate{err: errors.New("")}, returnFindByTitle{}, returnFindByID{})}

	result := sut.Create(&entities.BookDTO{
		Title:             "title",
		Author:            "author",
		PublishingCompany: "company",
		Edition:           1,
	})

	if result.StatusCode != 500 {
		t.Errorf("Should return status code %d, when is spected to be 500", result.StatusCode)
	}
}
