package applications

import (
	"restapi/src/applications/protocols"
	"restapi/src/shared"
)

type IGetOneBookUsecase interface {
	GetOne(id uint64) *shared.HTTPResponse
}
type getOneBookUsecase struct {
	repository protocols.IBooksRepository
}

func (u *getOneBookUsecase) GetOne(id uint64) *shared.HTTPResponse {
	entity, err := u.repository.FindByID(id)
	if err != nil {
		return shared.HTTPInternalServerError("Internal Server Error")
	}

	return shared.HTTPSuccess(entity)
}

func NewGetOneBookUsecase(repository protocols.IBooksRepository) IGetOneBookUsecase {
	return &getOneBookUsecase{repository: repository}
}
