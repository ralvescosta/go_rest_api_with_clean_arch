package applications

import (
	"restapi/src/applications/protocols"
	"restapi/src/shared"
)

// IGetAllBooksUsecase ...
type IGetAllBooksUsecase interface {
	GetAll() *shared.HTTPResponse
}
type getAllBooksUsecase struct {
	repository protocols.IBooksRepository
}

func (u *getAllBooksUsecase) GetAll() *shared.HTTPResponse {

	entities, err := u.repository.FindAll()
	if err != nil {
		return shared.HTTPInternalServerError("Internal Server Error")
	}

	return shared.HTTPSuccess(entities)
}

// NewGetAllBooksUsecase ...
func NewGetAllBooksUsecase(repository protocols.IBooksRepository) IGetAllBooksUsecase {
	return &getAllBooksUsecase{repository: repository}
}
