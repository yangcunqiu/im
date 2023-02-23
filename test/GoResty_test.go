package test

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/url"
	"testing"
)

type Attribution struct {
	Status  int    `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

func TestGet(t *testing.T) {
	client := resty.New()
	result := Attribution{}
	resp, err := client.R().
		SetResult(&result).
		SetQueryParam("shouji", "18865118837").
		SetAuthToken("APPCODE c09d6e05b220456a98633777f423a800").
		Get("https://jshmgsdmfb.market.alicloudapi.com/shouji/query")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp.StatusCode())
	fmt.Println(resp.Status())
	fmt.Println(resp.Header())
	fmt.Println(resp.Result())
}

func TestPost(t *testing.T) {
	client := resty.New()
	resp, err := client.R().
		SetBody(url.Values{"content": {"code:78945,expire_at:7"},
			"phone_number": {"18865118837"},
			"template_id":  {"TPL_0001"}}.Encode()).
		SetHeader("Authorization", "APPCODE c09d6e05b220456a98633777f423a800").
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		Post("https://dfsns.market.alicloudapi.com/data/send_sms")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp.StatusCode())
	fmt.Println(resp.Status())
	fmt.Println(resp.Header())
	fmt.Println(resp.Result())
	fmt.Println(resp.Request.Header)
	fmt.Println(resp.Request.Body)
	fmt.Println(resp.Request.URL)
}
