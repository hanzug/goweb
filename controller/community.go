package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"goweb/logic"
	"strconv"
)

func CommunityHandler(c *gin.Context) {
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList()", zap.Error(err))
		ResponseError(c, CodeServerBusy)
	}
	ResponseSuccess(c, data)
}

func CommunityDetailHandler(c *gin.Context) {
	communityID := c.Param("id")

	id, err := strconv.ParseInt(communityID, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	data, err := logic.GetCommunityDetail(id)
	if err != nil {
		zap.L().Error("logic.GetCommunityList()", zap.Error(err))
		ResponseErrorWithMsg(c, CodeServerBusy, id)
		return
	}
	ResponseSuccess(c, data)
}
