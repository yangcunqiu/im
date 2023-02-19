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

func TestEncode(t *testing.T) {
	str := ""
	salt := ""
	sha256 := utils.EncodeBySHA256(str, salt)
	fmt.Println(sha256)
}

func TestDecode(t *testing.T) {
	plaintext := ""
	sourceCiphertext := "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"
	salt := "1"
	ok := utils.VailPasswordBySHA256(plaintext, sourceCiphertext, salt)
	fmt.Println(ok)
}
