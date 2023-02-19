package service

import (
	"github.com/gin-gonic/gin"
	"im/dao"
	"im/handler"
	"im/utils"
)

func GetDistrictList(c *gin.Context) {
	list := dao.GetDistrictList()
	// 转成树形结构
	handler.Success(c, utils.ListToTree(list))
}
