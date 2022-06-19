package main

import (
	"context"
	"kitexdousheng/cmd/publish/service"
	"kitexdousheng/kitex_gen/publish"
	"kitexdousheng/pkg/errno"
)

// PublishSrvImpl implements the last service interface defined in the IDL.
type PublishSrvImpl struct{}

// PublishAction implements the PublishSrvImpl interface.
func (s *PublishSrvImpl) PublishAction(ctx context.Context, req *publish.DouyinPublishActionRequest) (resp *publish.DouyinPublishActionResponse, err error) {
	// TODO: Your code here...
	err = service.PublishAction(req.UserId, req.Title)
	if err != nil {
		return &publish.DouyinPublishActionResponse{
			StatusCode: errno.ServiceErr.ErrCode,
			StatusMsg:  &errno.ServiceErr.ErrMsg,
		}, nil
	}
	return &publish.DouyinPublishActionResponse{
		StatusCode: errno.Success.ErrCode,
		StatusMsg:  &errno.Success.ErrMsg,
	}, nil
}

// PublishList implements the PublishSrvImpl interface.
func (s *PublishSrvImpl) PublishList(ctx context.Context, req *publish.DouyinPublishListRequest) (resp *publish.DouyinPublishListResponse, err error) {
	// TODO: Your code here...
	videos, err := service.PublishList(req.UserId)
	if err != nil {
		return &publish.DouyinPublishListResponse{
			StatusCode: errno.ServiceErr.ErrCode,
			StatusMsg:  &errno.ServiceErr.ErrMsg,
		}, err
	}
	return &publish.DouyinPublishListResponse{
		StatusCode: errno.Success.ErrCode,
		StatusMsg:  &errno.Success.ErrMsg,
		VideoList:  videos,
	}, nil
}
