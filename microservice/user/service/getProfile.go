package service

import (
	"context"
	"workbench-user/dao"
	pb "workbench-user/proto"
	"workbench/pkg/errno"
)

// GetProfile ... 获取用户个人信息
func (s *UserService) GetProfile(ctx context.Context, req *pb.GetRequest, resp *pb.UserProfile) error {
	user, err := dao.GetUser(req.Id)
	if err != nil {
		return errno.ServerErr(errno.ErrDatabase, err.Error())
	}

	if user == nil {
		return errno.ServerErr(errno.ErrUserNotExisted, "")
	}
	resp.Id = user.Id
	resp.Name = user.Name
	resp.RealName = user.RealName
	resp.Avatar = user.Avatar
	resp.Email = user.Email
	resp.Tel = user.Tel
	resp.Role = user.Role
	resp.Team = user.TeamId
	resp.Group = user.GroupId

	return nil
}
