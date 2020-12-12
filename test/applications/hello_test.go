package test

import (
	"testing"

	hello "restapi/applications"
)

func TestHello(t *testing.T) {
	result := hello.Hello()

	if result != "Hello World" {
		t.Errorf("Hello return %s, want Hello World", result)
	}
}
