package applications

import (
	"restapi/src/applications/protocols"
	"restapi/src/entities"
	"restapi/src/shared"
	"time"
)

// CreateBookUsecase ...
type CreateBookUsecase struct {
	BooksRepository protocols.IBooksRepository
}

// Create ...
func (u *CreateBookUsecase) Create(body *entities.BookDTO) *shared.HTTPResponse {

	entity := &entities.BookEntity{
		Title:     body.Title,
		Author:    body.Author,
		Edition:   body.Edition,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err := u.BooksRepository.Create(entity)
	if err != nil {
		return shared.HTTPInternalServerError("Internal Server Error")
	}

	return shared.HTTPCreated()
}
