package applications

import "restapi/shared"

// CreateBookUsecase ...
type CreateBookUsecase struct{}

// Create ...
func (*CreateBookUsecase) Create(body interface{}) *shared.HTTPResponse {
	return shared.HTTPCreated()
}
