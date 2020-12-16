package applications

import (
	"restapi/src/shared"
	"time"
)

// IHealthUsecase ...
type IHealthUsecase interface {
	Health() *shared.HTTPResponse
}
type healthUsecase struct {
	startupTime time.Time
}

// Health ...
func (u *healthUsecase) Health() *shared.HTTPResponse {
	var response = make(map[string]interface{})

	response["startup"] = u.startupTime
	response["now"] = time.Now()

	return shared.HTTPSuccess(response)
}

// NewHealthUsecase ...
func NewHealthUsecase(startupTime time.Time) IHealthUsecase {
	return &healthUsecase{startupTime: startupTime}
}
