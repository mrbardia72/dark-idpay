package services

import (
	"../config"
	"../helper"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)


var logpayCollection = config.DbConfig().Database("idpay").Collection("logpay")
var keypayCollection = config.DbConfig().Database("idpay").Collection("keypay")

func CreateTNX()  {
	ctx := context.Background()
	date_now := time.Now().Format("02-01-2006")
	time_now := time.Now().Format("15:04:05")

	data := map[string]string{
		"order_id": "444",
		"amount":   "23000",
		"name":     "بردیا کاظمی",
		"phone":    "09360750299",
		"mail":     "langroud@gilan.iran",
		"desc":     "توضیحات پرداخت کننده",
		"callback": "https://mrb.com/v/lg/34ew",
		"date-now":	date_now,
		"time-now":	time_now,
	}
	get_order_id := data["order_id"]
	payload, _ := json.Marshal(data)
	req, _ := http.NewRequest("POST", config.CREATE_URL_TNX, bytes.NewBuffer(payload))
	ctx1,_ :=context.WithTimeout(ctx,time.Minute)
	req=req.WithContext(ctx1)
	helper.ReqHeader(req)

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	app1 := helper.UnJSON(body)

	insertResult, err := logpayCollection.InsertOne(ctx, &data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("create order: ", insertResult)

	data1 := map[string]string{
		"order_id": get_order_id,
		"key":   	app1.Id,
		"url":   	app1.Link,
		"date-now":	date_now,
		"time-now":	time_now,
	}

	// create key for create tnx
	insertResult1, err := keypayCollection.InsertOne(ctx, &data1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("create key order: ", insertResult1)
}


