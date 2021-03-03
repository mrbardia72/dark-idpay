package helper

import (
	"../models"
	"encoding/json"
)

func UnJSON(body []byte) models.TnxRcv1 {
	var app models.TnxRcv1
	_ = json.Unmarshal(body, &app)
	return app
}