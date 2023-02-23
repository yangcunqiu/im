package test

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"testing"
)

func TestHttpGet(t *testing.T) {
	appCode := "c09d6e05b220456a98633777f423a800"
	targetUrl := "https://jshmgsdmfb.market.alicloudapi.com/shouji/query"
	phone := "18865118837"

	request, _ := http.NewRequest("GET", targetUrl+"?shouji="+phone, nil)
	request.Header.Add("Authorization", "APPCODE "+appCode)
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Println(err)
		return
	}
	defer response.Body.Close()
	all, _ := io.ReadAll(response.Body)
	fmt.Println(string(all))
}

func TestHttpPost(t *testing.T) {
	appCode := "c09d6e05b220456a98633777f423a800"
	targetUrl := "https://dfsns.market.alicloudapi.com/data/send_sms"
	phone := "18865118837"
	code := "145236"
	expire := 5
	templateId := "TPL_0001"
	values := url.Values{
		"content":      {"code:" + code + ",expire_at:" + strconv.Itoa(expire)},
		"phone_number": {phone},
		"template_id":  {templateId},
	}

	request, _ := http.NewRequest("POST", targetUrl, strings.NewReader(values.Encode()))
	request.Header.Add("Authorization", "APPCODE "+appCode)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Println(err)
		return
	}
	defer response.Body.Close()
	all, _ := io.ReadAll(response.Body)
	fmt.Println(string(all))
}
