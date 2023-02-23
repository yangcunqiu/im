package invoke

import (
	"encoding/json"
	"im/dao"
	"im/global"
	"im/model"
	"im/model/third"
	"net/url"
	"time"
)

var phoneCodeServiceType = "发送手机验证码"

func QueryPhoneAttribution() {

}

func SendVerifyCode(phone string, code string, expire int) bool {
	startTime := time.Now()
	var codeResult third.PhoneVerify
	resp, err := global.HttpClient.R().SetResult(&codeResult).SetBody(url.Values{
		"content":      {"code:" + code + ",expire_at:2"},
		"phone_number": {phone},
		"template_id":  {global.Config.Third.Aliyun.PhoneSendVerifyCodeTemplateId},
	}.Encode()).SetHeaders(map[string]string{
		"Authorization": "APPCODE " + global.Config.Third.Aliyun.AppCode,
		"Content-Type":  "application/x-www-form-urlencoded",
	}).Post(global.Config.Third.Aliyun.PhoneSendVerifyCodeUrl)
	endTime := time.Now()
	cost := endTime.Sub(startTime).Milliseconds()

	if err != nil {
		// 记录日志
		callLog := model.CallLog{
			Type:          1,
			Url:           global.Config.Third.Aliyun.IPAttributionUrl,
			MethodType:    "GET",
			ServiceType:   phoneCodeServiceType,
			InvokeStatus:  -1,
			ServiceStatus: "-1",
			ErrorString:   err.Error(),
			RequestTime:   startTime,
			ResponseTime:  startTime,
			Cost:          cost,
		}
		dao.AddCallLog(callLog)
		return false
	}

	marshal, _ := json.Marshal(codeResult)
	callLog := model.CallLog{
		Type:          1,
		Url:           global.Config.Third.Aliyun.IPAttributionUrl,
		MethodType:    "GET",
		ServiceType:   phoneCodeServiceType,
		InvokeStatus:  resp.StatusCode(),
		ServiceStatus: codeResult.Status,
		RequestStr:    resp.Request.QueryParam.Encode(),
		RequestTime:   startTime,
		ResponseStr:   string(marshal),
		ResponseTime:  endTime,
		Cost:          cost,
	}
	dao.AddCallLog(callLog)

	if resp.StatusCode() == 200 {
		if codeResult.Status == "OK" {
			// 成功
			return true
		}
	}
	return false
}
