package handler

import (
	"github.com/gin-gonic/gin"
	"im/model"
	"strconv"
)

func PageOf(pageNumber, pageSize int, total int64, list any) model.Page {
	return model.Page{
		PageNum:  pageNumber,
		PageSize: pageSize,
		Total:    total,
		List:     list,
	}
}

func PageParams(c *gin.Context) (pageNum, pageSize, offset int) {
	pageNum, _ = strconv.Atoi(c.DefaultQuery("pageNum", "1"))
	pageSize, _ = strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	if pageSize > 1000 {
		pageSize = 1000
	}
	offset = (pageNum - 1) * pageSize
	return
}
