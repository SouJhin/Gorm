package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"server/models"
)

// GetProblemList
// @Tags 获取问题列表
// @Summary 传入食物名称和价格
// @Param name query string false "请输入食物名称"
// @Param size query int false "请输入食物价格"
// Success 200 {string} json "{"code":"200", "msg", "data":""}"
// @Router /addFood [get]
func GetProblemList(ctx *gin.Context) {
	models.GetProblemList()
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
	})
}
