package user

import (
	"context"

	. "workbench-gateway/handler"
	"workbench-gateway/service"
	"workbench-gateway/util"
	pb "workbench-user/proto"
	"workbench/log"
	"workbench/pkg/errno"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// UpdateInfo ... 修改用户个人信息
// @Summary update info api
// @Description 修改用户个人信息
// @Tags user
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "token 用户令牌"
// @Param object body UpdateInfoRequest  true "update_info_request"
// @Success 200 {object} handler.Response
// @Router /user [put]
func UpdateInfo(c *gin.Context) {
	log.Info("User getInfo function called.", zap.String("X-Request-Id", util.GetReqID(c)))

	var req UpdateInfoRequest
	if err := c.BindJSON(&req); err != nil {
		SendError(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	userId := c.MustGet("userID").(uint32)

	// 构造请求给 getInfo
	updateInfoReq := &pb.UpdateInfoRequest{
		Id: userId,
		Info: &pb.UserInfo{
			Name:      req.Name,
			RealName:  req.RealName,
			AvatarUrl: req.AvatarURL,
			Email:     req.Email,
			Tel:       req.Tel,
		},
	}

	// 发送请求
	_, err := service.UserClient.UpdateInfo(context.TODO(), updateInfoReq)
	if err != nil {
		SendError(c, err, nil, "", GetLine())
		return
	}

	SendResponse(c, nil, nil)
}
