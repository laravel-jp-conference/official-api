package foundation

import (
	"net/http"
	"net/url"
)

// RequestAction Request時の値操作
type RequestAction struct {
	*http.Request
}

// ResolveRequestURI リクエスト時のURI取得 vnd.error利用時に返却
func (a *RequestAction) ResolveRequestURI() (uv *url.URL) {
	uv, _ = url.Parse(a.URL.String())
	uv.Scheme = string("https")
	uv.Path = a.URL.Path
	uv.Host = a.Host
	uv.RawQuery = ""
	return
}
