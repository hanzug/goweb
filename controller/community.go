package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"goweb/logic"
)

func CommunityHandler(c *gin.Context) {
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList()", zap.Error(err))
		ResponseError(c, CodeServerBusy)
	}
	ResponseSuccess(c, data)
}
