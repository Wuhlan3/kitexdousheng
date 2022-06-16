package main

import (
	"context"
	"kitexdousheng/kitex_gen/feed"
)

// FeedSrvImpl implements the last service interface defined in the IDL.
type FeedSrvImpl struct{}

// GetUserFeed implements the FeedSrvImpl interface.
func (s *FeedSrvImpl) GetUserFeed(ctx context.Context, req *feed.DouyinFeedRequest) (resp *feed.DouyinFeedResponse, err error) {
	// TODO: Your code here...
	return
}

// GetVideoById implements the FeedSrvImpl interface.
func (s *FeedSrvImpl) GetVideoById(ctx context.Context, req *feed.VideoIdRequest) (resp *feed.Video, err error) {
	// TODO: Your code here...
	return
}
