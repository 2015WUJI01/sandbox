package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const limit = 5
const skip = 0

type _ProductInfoList []string

type _ProductInfoName []struct {
	Typename string `json:"__typename"`
	Name     string `json:"name"`
	Value    string `json:"value"`
}

type _ProductInfoTitle [][][]__Title

type __Title struct {
	CustomName  string      `json:"customName"`
	Description interface{} `json:"description"`
	IsAdminOnly bool        `json:"isAdminOnly"`
	Section     interface{} `json:"section"`
	Title       string      `json:"title"`
	UnitString  interface{} `json:"unitString"`
	Units       interface{} `json:"units"`
	Value       string      `json:"value"`
}

type T struct {
	Datas []struct {
		Id struct {
			Oid string `json:"$oid"`
		} `json:"_id"`
		Brand         string      `json:"brand"`
		Category      string      `json:"category"`
		Condition     string      `json:"condition"`
		CouponCode    string      `json:"coupon_code"`
		CreatedAt     string      `json:"created_at"`
		Currency      string      `json:"currency"`
		CurrentPrice  interface{} `json:"current_price"`
		Description   string      `json:"description"`
		Discount      interface{} `json:"discount"`
		Image         interface{} `json:"image"`
		Keyword       string      `json:"keyword"`
		OutOfStock    interface{} `json:"out_of_stock"`
		Platform      string      `json:"platform"`
		PreviousPrice interface{} `json:"previous_price"`
		ProductId     string      `json:"product_id"`
		ProductInfo   interface{} `json:"product_info"`
		ProductLink   string      `json:"product_link"`
		Stars         interface{} `json:"stars"`
		SubCategory   string      `json:"sub_category"`
		Title         string      `json:"title"`
		UpdatedAt     string      `json:"updated_at"`
	} `json:"datas"`
	DocsCount int         `json:"docs_count"`
	EndTime   interface{} `json:"end_time"`
	StartTime interface{} `json:"start_time"`
	UsedTime  float64     `json:"used_time"`
}

func main() {
	var err error

	// 发送请求
	resp, err := http.Get(fmt.Sprintf("http://172.105.11.118/api/get_deal_info?collection_name=afc_product&limit=%d&skip=%d", limit, skip))
	if err != nil {
		panic(err)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	t := T{}
	err = json.Unmarshal(b, &t)
	if err != nil {
		panic(err)
	}

	for _, data := range t.Datas {
		switch data.ProductInfo.(type) {
		case _ProductInfoList:
			log.Printf("list")
		case _ProductInfoTitle:
			log.Printf("list")
		case _ProductInfoName:
			log.Printf("name")
		case []string:
			log.Printf("[]string")
		case [][][]interface{}:
			log.Printf("[][][]interface{}")
		case []map[string]string:
			log.Printf("[]map[string]string")
		case string:
			log.Printf("string")
		default:
			log.Printf("unknown")
		}
		fmt.Println()
	}

}
