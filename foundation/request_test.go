package foundation

import (
	"net/http"
	"net/url"
	"testing"
)

func TestRequestAction_ResolveRequestURI(t *testing.T) {
	ra := &RequestAction{
		&http.Request{
			URL: &url.URL{
				Scheme:     "https",
				Host:       "conference2019.laravel.jp",
				Path:       "testing",
				ForceQuery: false,
			},
			Host: "conference2019.laravel.jp",
		},
	}
	res := ra.ResolveRequestURI().String()
	if res != "https://conference2019.laravel.jp/testing" {
		t.Fatalf("it is not expected value %#v", res)
	}
}
