package service

import (
	"context"
	"workbench-user/dao"

	pb "workbench-user/proto"
	"workbench/pkg/errno"
)

// UpdateInfo ... 更新用户信息
func (s *UserService) UpdateInfo(ctx context.Context, req *pb.UpdateInfoRequest, resp *pb.Response) error {
	user, err := dao.GetUser(req.Id)
	if err != nil {
		return errno.ServerErr(errno.ErrDatabase, err.Error())
	}

	if user == nil {
		return errno.ServerErr(errno.ErrUserNotExisted, err.Error())
	}

	user.Name = req.Info.Name
	user.RealName = req.Info.RealName
	user.Avatar = req.Info.AvatarUrl
	user.Tel = req.Info.Tel

	if err := user.Save(); err != nil {
		return errno.ServerErr(errno.ErrDatabase, err.Error())
	}

	return nil
}
