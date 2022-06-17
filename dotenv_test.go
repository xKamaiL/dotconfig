package dotconfig

import "testing"

type Spec struct {
	ApiUrl string `required:"true"`
}

func TestLoad(t *testing.T) {
	var sp Spec
	err := Load(&sp, "./testdata/.env")
	if err != nil {
		t.Error(err)
		return
	}
	if sp.ApiUrl != "https://domain.com" {
		t.Error("ApiUrl", sp.ApiUrl)
		return
	}
}
