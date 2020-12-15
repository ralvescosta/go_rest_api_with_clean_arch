package applications

import (
	"restapi/src/applications/protocols"
	"restapi/src/entities"
	"restapi/src/shared"
	"time"
)

// ICreateBookUsecase ...
type ICreateBookUsecase interface {
	Create(bookDTO *entities.BookDTO) *shared.HTTPResponse
}

type createBookUsecase struct {
	repository protocols.IBooksRepository
}

// Create ...
func (u *createBookUsecase) Create(bookDTO *entities.BookDTO) *shared.HTTPResponse {

	entity, err := u.repository.FindByTitle(bookDTO.Title)
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

	_, err = u.repository.Create(entity)
	if err != nil {
		return shared.HTTPInternalServerError("Internal Server Error")
	}

	return shared.HTTPCreated()
}

// NewCreateBookUsecase ...
func NewCreateBookUsecase(repository protocols.IBooksRepository) ICreateBookUsecase {
	return &createBookUsecase{repository: repository}
}
