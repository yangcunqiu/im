package test

import (
	"fmt"
	"im/utils"
	"testing"
)

func TestGetOS(t *testing.T) {
	//userAgent := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"
	userAgent := "PostmanRuntime/7.31.0"
	os := utils.GetOSVersionByUserAgent(userAgent)
	fmt.Println(os)
}

func TestGetBrowser(t *testing.T) {
	userAgent := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) /109.0.0.0 Safari/537.36"
	browser := utils.GetBrowserByUserAgent(userAgent)
	fmt.Println(browser)
}
