package protocols

import "restapi/src/entities"

// IBooksRepository ...
type IBooksRepository interface {
	Create(entity *entities.BookEntity) (*entities.BookEntity, error)
	FindByTitle(title string) (*entities.BookEntity, error)
	FindByID(id uint64) (*entities.BookEntity, error)
	FindAll() ([]entities.BookEntity, error)
	Delete(id uint64) (*entities.BookEntity, error)
}
