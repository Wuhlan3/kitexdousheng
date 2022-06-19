package main

import (
	"context"
	"kitexdousheng/cmd/relation/service"
	"kitexdousheng/kitex_gen/relation"
	"kitexdousheng/pkg/errno"
	"log"
	"strconv"
)

// RelationSrvImpl implements the last service interface defined in the IDL.
type RelationSrvImpl struct{}

// RelationAction implements the RelationSrvImpl interface.
func (s *RelationSrvImpl) RelationAction(ctx context.Context, req *relation.DouyinRelationActionRequest) (resp *relation.DouyinRelationActionResponse, err error) {
	log.Println(req)

	err = service.RelationAction(req.MyId, req.HisId, strconv.Itoa(int(req.ActionType)))
	if err != nil {
		return &relation.DouyinRelationActionResponse{
			StatusCode: errno.ServiceErr.ErrCode,
			StatusMsg:  &errno.ServiceErr.ErrMsg,
		}, err
	}
	return &relation.DouyinRelationActionResponse{
		StatusCode: 0,
		StatusMsg:  &errno.Success.ErrMsg,
	}, err

}

// RelationFollowList implements the RelationSrvImpl interface.
func (s *RelationSrvImpl) RelationFollowList(ctx context.Context, req *relation.DouyinRelationFollowListRequest) (*relation.DouyinRelationFollowListResponse, error) {
	followList, err := service.RelationFollowList(req.UserId)
	if err != nil {
		return &relation.DouyinRelationFollowListResponse{
			StatusCode: errno.ServiceErr.ErrCode,
			StatusMsg:  &errno.ServiceErr.ErrMsg,
			UserList:   nil,
		}, err
	}
	return &relation.DouyinRelationFollowListResponse{
		StatusCode: 0,
		StatusMsg:  &errno.Success.ErrMsg,
		UserList:   followList,
	}, err
}

// RelationFollowerList implements the RelationSrvImpl interface.
func (s *RelationSrvImpl) RelationFollowerList(ctx context.Context, req *relation.DouyinRelationFollowerListRequest) (*relation.DouyinRelationFollowerListResponse, error) {
	followerList, err := service.RelationFollowerList(req.UserId)
	if err != nil {
		return &relation.DouyinRelationFollowerListResponse{
			StatusCode: errno.ServiceErr.ErrCode,
			StatusMsg:  &errno.ServiceErr.ErrMsg,
			UserList:   nil,
		}, err
	}
	return &relation.DouyinRelationFollowerListResponse{
		StatusCode: 0,
		StatusMsg:  &errno.Success.ErrMsg,
		UserList:   followerList,
	}, err
}
