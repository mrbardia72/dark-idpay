package services

import (
	"../config"
	"../helper"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"context"
)

func InquiryPayment()  {
	ctx:=context.Background()
	data := map[string]string{
		"id":       "d2e353189823079e1e4181772cff5292",
		"order_id": "101",
	}

	payload, _ := json.Marshal(data)

	req, _ := http.NewRequest("POST", config.INQUIRY_URL, bytes.NewBuffer(payload))
	ctx1,_ :=context.WithTimeout(ctx,time.Millisecond*100)
	req=req.WithContext(ctx1)
	helper.ReqHeader(req)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}