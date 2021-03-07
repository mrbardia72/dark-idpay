package services

import (
	"../config"
	"../helper"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func TransactionsPayment() {
	ctx:=context.Background()
	data := map[string]string{
		"id": "e22952579725883bbad9f8fa429134bf",
		"order_id": "101",
		"amount": "10000",
		"status": "100",
		"track_id": "27384837",
		"payment_card_no": "636214******5409",
		"payment_hashed_card_no": "B913D97F01CE42601181135DF3D0F81DA9E98E61BE3E3AB4436E6345D6AB0AEA",
		//"payment_date": {"min": 1600005000, "max": 1600006000},
		//"settlement_date": {"min": 1600005746, "max": 1600006000}
	}

	payload, _ := json.Marshal(data)

	req, _ := http.NewRequest("POST", config.TRANSACTIONS_URL, bytes.NewBuffer(payload))
	ctx1,_ :=context.WithTimeout(ctx,time.Millisecond*100)
	req=req.WithContext(ctx1)
	helper.ReqHeader(req)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}
