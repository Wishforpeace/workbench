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

// Register ... 注册
// @Summary register api
// @Description register user
// @Tags auth
// @Accept application/json
// @Produce application/json
// @Param object body RegisterRequest true "register_request"
// @Success 200 {object} handler.Response
// @Router /auth/signup [post]
func Register(c *gin.Context) {
	log.Info("User register function called.", zap.String("X-Request-Id", util.GetReqID(c)))

	// 从前端获取请求
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		SendError(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// 构造请求给 register
	registerReq := &pb.RegisterRequest{
		Email:    req.Email,
		Name:     req.Name,
		Password: req.Password,
	}

	// 注册
	_, err := service.UserClient.Register(context.TODO(), registerReq)
	if err != nil {
		SendError(c, err, nil, "", GetLine())
		return
	}

	SendResponse(c, nil, nil)
}
