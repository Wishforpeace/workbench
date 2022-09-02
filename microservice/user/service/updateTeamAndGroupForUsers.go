package service

import (
	"context"
	"workbench-user/dao"

	pb "workbench-user/proto"
	"workbench/pkg/errno"
)

// UpdateTeamAndGroupForUsers update user's team or group.
func (s *UserService) UpdateTeamAndGroupForUsers(ctx context.Context, req *pb.UpdateTeamGroupRequest, resp *pb.Response) error {
	if err := dao.UpdateTeamAndGroup(req.Ids, req.Value, req.Kind); err != nil {
		return errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	return nil
}
