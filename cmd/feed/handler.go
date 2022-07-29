package main

import (
	"context"
	"kitexdousheng/cmd/feed/service"
	"kitexdousheng/kitex_gen/feed"
	"kitexdousheng/pkg/errno"
	"time"
)

// FeedSrvImpl implements the last service interface defined in the IDL.
type FeedSrvImpl struct{}

// Feed implements the FeedSrvImpl interface.
func (s *FeedSrvImpl) Feed(ctx context.Context, req *feed.DouyinFeedRequest) (resp *feed.DouyinFeedResponse, err error) {
	// TODO: Your code here...
	videos, err := service.Feed(req.UserId, *req.LatestTime)
	if err != nil {
		return &feed.DouyinFeedResponse{
			StatusCode: errno.ServiceErr.ErrCode,
			StatusMsg:  &errno.ServiceErr.ErrMsg,
		}, nil
	}
	//视频加载没有出错
	nextTime := time.Now().Unix()
	return &feed.DouyinFeedResponse{
		StatusCode: errno.Success.ErrCode,
		StatusMsg:  &errno.Success.ErrMsg,
		VideoList:  videos,
		NextTime:   &nextTime,
	}, nil

}
