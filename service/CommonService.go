package service

import (
	"github.com/gin-gonic/gin"
	"im/dao"
	"im/handler"
	"im/utils"
	"strconv"
)

func GetDistrictList(c *gin.Context) {
	pid, err := strconv.Atoi(c.DefaultQuery("pid", "0"))
	if err != nil {
		handler.Fail(c, handler.ParamsBindingError, "参数类型不匹配")
	}

	// 查全部
	list := dao.GetDistrictList()

	// 转成树形结构
	handler.Success(c, utils.ListToTree(list, pid))
}
