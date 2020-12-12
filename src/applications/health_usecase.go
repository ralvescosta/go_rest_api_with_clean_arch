package applications

import (
	"restapi/shared"
	"time"
)

// IHealthUsecase ...
type IHealthUsecase interface {
	Health() (*shared.HTTPResponse, error)
}

type healthUsecase struct {
	startupTime time.Time
}

// Health ...
func (u *healthUsecase) Health() (*shared.HTTPResponse, error) {
	var response = make(map[string]interface{})

	response["startup"] = u.startupTime
	response["now"] = time.Now()

	return shared.HTTPSuccess(response), nil
}

// NewHealthUsecase ...
func NewHealthUsecase(startupTime time.Time) IHealthUsecase {
	return &healthUsecase{startupTime: startupTime}
}
