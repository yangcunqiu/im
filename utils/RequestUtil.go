package utils

import "strings"

func GetOSVersionByUserAgent(userAgent string) string {
	// Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36
	var os string
	userAgentLower := strings.ToLower(userAgent)
	if strings.Contains(userAgentLower, "windows") {
		os = "Windows"
	} else if strings.Contains(userAgentLower, "mac") {
		os = "Mac"
	} else if strings.Contains(userAgentLower, "x11") {
		os = "Unix"
	} else if strings.Contains(userAgentLower, "android") {
		os = "Android"
	} else if strings.Contains(userAgentLower, "iphone") {
		os = "iOS"
	} else if strings.Contains(userAgentLower, "linux") {
		os = "Linux"
	} else if strings.Contains(userAgentLower, "postman") {
		os = "Postman"
	}

	// 保留原始信息
	var source string
	index := strings.Index(userAgent, "(")
	lastIndex := strings.Index(userAgent, ")")
	if index != -1 && lastIndex != -1 {
		source = userAgent[index:lastIndex] + ")"
	}
	return os + " " + source
}

func GetBrowserByUserAgent(userAgent string) string {
	// Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36
	var browser string
	userAgentLower := strings.ToLower(userAgent)
	if strings.Contains(userAgentLower, "chrome") && !strings.Contains(userAgentLower, "edg") {
		browser = "Chrome"
	} else if strings.Contains(userAgentLower, "firefox") {
		browser = "Firefox"
	} else if strings.Contains(userAgentLower, "safari") && !strings.Contains(userAgentLower, "chrome") && !strings.Contains(userAgentLower, "edg") {
		browser = "Safari"
	} else if strings.Contains(userAgentLower, "edg") {
		browser = "Edge"
	}
	return browser
}
