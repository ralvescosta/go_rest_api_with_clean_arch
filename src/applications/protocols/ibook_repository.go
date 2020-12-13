package protocols

import "restapi/src/entities"

// IBooksRepository ...
type IBooksRepository interface {
	Create(entity *entities.BookEntity) (*entities.BookEntity, error)
}
