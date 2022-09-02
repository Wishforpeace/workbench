package user

import (
	"context"
	// "strconv"

	. "workbench-gateway/handler"
	pb "workbench-user/proto"
	"workbench/log"
	"workbench/pkg/errno"

	"go.uber.org/zap"

	// "workbench-gateway/pkg/token"
	"workbench-gateway/service"
	"workbench-gateway/util"

	"github.com/gin-gonic/gin"
)

// UpdateTeamAndGroupForUsers ... 通过 teamid 或 groupid 给 users 数组分组/团队
func UpdateTeamAndGroupForUsers(c *gin.Context) {
	log.Info("User getInfo function called.", zap.String("X-Request-Id", util.GetReqID(c)))

	// 从前端获取请求
	var req UpdateTeamGroupRequest
	if err := c.Bind(&req); err != nil {
		SendError(c, errno.ErrBind, nil, err.Error(), GetLine())
		return
	}

	// 构造请求给 updateTeamGroup
	updateTeamGroupReq := &pb.UpdateTeamGroupRequest{
		Value: req.Value,
		Kind:  req.Kind,
	}
	for i := 0; i < len(req.Ids); i++ {
		updateTeamGroupReq.Ids = append(updateTeamGroupReq.Ids, req.Ids[i])
	}

	// 发送请求
	_, err := service.UserClient.UpdateTeamAndGroupForUsers(context.TODO(), updateTeamGroupReq)
	if err != nil {
		SendError(c, err, nil, "", GetLine())
		return
	}

	SendResponse(c, nil, nil)
}
