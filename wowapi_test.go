package wowapi_test

import (
	"os"
	"testing"

	"github.com/raph6/wowapi"
)

var req wowapi.RequestFunc

func TestMain(m *testing.M) {
	// env var
	token := os.Getenv("API_CLIENT_ID")
	secret := os.Getenv("API_SECRET")
	var err error
	req, err = wowapi.Client(token, secret)
	if err != nil {
		panic(err)
	}

	code := m.Run()

	os.Exit(code)
}
