package applications

import (
	"testing"
	"time"
)

func TestHealthUsecase(t *testing.T) {
	startupMocked := time.Now()
	healthUsecase := HealthUsecase{StartupTime: startupMocked}

	httpResponse := healthUsecase.Health()
	body := httpResponse.Body.(map[string]interface{})

	if httpResponse.StatusCode != 200 && body["startup"] != startupMocked {
		t.Errorf("Should return api health")
	}
}
