package applications

import (
	"fmt"
	"restapi/entities"
	"restapi/shared"
	"time"
)

// CreateBookUsecase ...
type CreateBookUsecase struct{}

// Create ...
func (*CreateBookUsecase) Create(body *entities.BookDTO) *shared.HTTPResponse {

	entity := entities.BookEntity{
		Title:     body.Title,
		Author:    body.Author,
		Edition:   body.Edition,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	fmt.Println(entity)

	return shared.HTTPCreated()
}
