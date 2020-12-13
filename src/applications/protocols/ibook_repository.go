package protocols

import "restapi/src/entities"

// IBooksRepository ...
type IBooksRepository interface {
	Create(book *entities.BookEntity) (*entities.BookEntity, error)
}
