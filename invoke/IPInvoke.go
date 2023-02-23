package invoke

import (
	"encoding/json"
	"fmt"
	"im/dao"
	"im/global"
	"im/model"
	"im/model/third"
	"time"
)

var IPServiceType = "IP归属地查询"

func QueryIPAttribution(ip string) string {
	startTime := time.Now()
	var ipAttr third.IPAttribution
	resp, err := global.HttpClient.R().
		SetResult(&ipAttr).
		SetQueryParam("ip", ip).
		SetHeader("Authorization", "APPCODE "+global.Config.Third.Aliyun.AppCode).
		Get(global.Config.Third.Aliyun.IPAttributionUrl)
	endTime := time.Now()
	cost := endTime.Sub(startTime).Milliseconds()

	if err != nil {
		// 记录日志
		callLog := model.CallLog{
			Type:          1,
			Url:           global.Config.Third.Aliyun.IPAttributionUrl,
			MethodType:    "GET",
			ServiceType:   IPServiceType,
			InvokeStatus:  -1,
			ServiceStatus: "-1",
			ErrorString:   err.Error(),
			RequestTime:   startTime,
			ResponseTime:  startTime,
			Cost:          cost,
		}
		dao.AddCallLog(callLog)
		return ""
	}

	// 记录日志
	marshal, _ := json.Marshal(ipAttr)
	callLog := model.CallLog{
		Type:          1,
		Url:           global.Config.Third.Aliyun.IPAttributionUrl,
		MethodType:    "GET",
		ServiceType:   IPServiceType,
		InvokeStatus:  resp.StatusCode(),
		ServiceStatus: string(rune(ipAttr.Status)),
		RequestStr:    resp.Request.QueryParam.Encode(),
		RequestTime:   startTime,
		ResponseStr:   string(marshal),
		ResponseTime:  endTime,
		Cost:          cost,
	}
	dao.AddCallLog(callLog)

	if resp.StatusCode() == 200 {
		if ipAttr.Status == 0 {
			// 成功
			return fmt.Sprintf("%v-%v-%v-%v (lat:%v, lng:%v)", ipAttr.Result.AdInfo.Nation, ipAttr.Result.AdInfo.Province,
				ipAttr.Result.AdInfo.City, ipAttr.Result.AdInfo.District, ipAttr.Result.Location.Lat, ipAttr.Result.Location.Lng)
		} else if ipAttr.Status == 375 {
			return "局域网IP"
		}
	}
	return ""
}
