package service

import (
	"context"
	"workbench-user/dao"

	pb "workbench-user/proto"
	"workbench/pkg/errno"
)

// List ... 获取用户列表
func (s *UserService) List(ctx context.Context, req *pb.ListRequest, resp *pb.ListResponse) error {

	// 过滤条件
	filter := &dao.UserModel{TeamId: req.Team}
	if req.Group != 0 {
		filter.GroupId = req.Group
	}

	// DB 查询
	list, err := dao.ListUser(req.Offset, req.Limit, req.LastId, filter)
	if err != nil {
		return errno.ServerErr(errno.ErrDatabase, err.Error())
	}

	resList := make([]*pb.User, len(list))

	for i, item := range list {
		resList[i] = &pb.User{
			Id:       item.Id,
			Name:     item.Name,
			RealName: item.RealName,
			Email:    item.Email,
			Avatar:   item.Avatar,
			Role:     item.Role,
			Team:     item.TeamId,
			Group:    item.GroupId,
		}
	}

	resp.Count = uint32(len(list))
	resp.List = resList

	return nil
}
