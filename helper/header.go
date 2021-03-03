package helper

import (
	"../config"
	"net/http"
)

func ReqHeader(req *http.Request) {
	req.Header.Set("Content-Type", config.CONTENT_TYPE)
	req.Header.Set("X-API-KEY", config.X_API_KEY)
	req.Header.Set("X-SANDBOX", config.X_SANDBOX)
}