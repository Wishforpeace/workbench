package service

import (
	"context"

	"workbench-user/dao"
	pb "workbench-user/proto"
	"workbench/pkg/errno"
)

// GetInfo ... 获取用户信息
func (s *UserService) GetInfo(ctx context.Context, req *pb.GetInfoRequest, resp *pb.UserInfoResponse) error {

	list, err := dao.GetUserByIds(req.Ids)
	if err != nil {
		return errno.ServerErr(errno.ErrDatabase, err.Error())
	}

	userInfos := make([]*pb.UserInfo, len(list))

	for i, user := range list {
		userInfos[i] = &pb.UserInfo{
			Id:        user.Id,
			Name:      user.Name,
			RealName:  user.RealName,
			AvatarUrl: user.Avatar,
			Email:     user.Email,
		}
	}

	resp.List = userInfos

	return nil
}
