package mocks

import (
	"restapi/src/applications/protocols"
	"restapi/src/entities"
)

type inMemoryRepository struct {
	returnCreate      ReturnCreate
	returnFindByTitle ReturnFindByTitle
	returnFindByID    ReturnFindByID
	returnFindAll     ReturnFindAll
	returnDelete      ReturnDelete
}

// ReturnCreate ...
type ReturnCreate struct {
	Entity *entities.BookEntity
	Err    error
}

// ReturnFindByTitle ...
type ReturnFindByTitle struct {
	Entity *entities.BookEntity
	Err    error
}

// ReturnFindByID ...
type ReturnFindByID struct {
	Entity *entities.BookEntity
	Err    error
}

// ReturnFindAll ...
type ReturnFindAll struct {
	Entity []entities.BookEntity
	Err    error
}

// ReturnDelete ...
type ReturnDelete struct {
	Entity *entities.BookEntity
	Err    error
}

func (repo *inMemoryRepository) Create(entity *entities.BookEntity) (*entities.BookEntity, error) {
	return repo.returnCreate.Entity, repo.returnCreate.Err
}
func (repo *inMemoryRepository) FindByTitle(title string) (*entities.BookEntity, error) {
	return repo.returnFindByTitle.Entity, repo.returnFindByTitle.Err
}
func (repo *inMemoryRepository) FindByID(id uint64) (*entities.BookEntity, error) {
	return repo.returnFindByID.Entity, repo.returnFindByID.Err
}
func (repo *inMemoryRepository) FindAll() ([]entities.BookEntity, error) {
	return repo.returnFindAll.Entity, repo.returnFindAll.Err
}
func (repo *inMemoryRepository) Delete(id uint64) (*entities.BookEntity, error) {
	return repo.returnDelete.Entity, repo.returnFindAll.Err
}

// NewInMemoryRepository ...
func NewInMemoryRepository(returnCreate ReturnCreate, returnFindByTitle ReturnFindByTitle, returnFindByID ReturnFindByID, returnFindAll ReturnFindAll, returnDelete ReturnDelete) protocols.IBooksRepository {
	return &inMemoryRepository{returnCreate: returnCreate, returnFindByTitle: returnFindByTitle, returnFindByID: returnFindByID, returnFindAll: returnFindAll, returnDelete: returnDelete}
}
