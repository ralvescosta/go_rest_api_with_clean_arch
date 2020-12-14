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
func (u *CreateBookUsecase) Create(bookDTO *entities.BookDTO) *shared.HTTPResponse {

	entity, err := u.BooksRepository.FindByTitle(bookDTO.Title)
	if err != nil {
		return shared.HTTPInternalServerError("Internal Server Error")
	}
	if entity != nil {
		return shared.HTTPConflict("Book title already registered")
	}

	entity = &entities.BookEntity{
		Title:             bookDTO.Title,
		Author:            bookDTO.Author,
		Edition:           bookDTO.Edition,
		PublishingCompany: bookDTO.PublishingCompany,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}

	_, err = u.BooksRepository.Create(entity)
	if err != nil {
		return shared.HTTPInternalServerError("Internal Server Error")
	}

	return shared.HTTPCreated()
}
