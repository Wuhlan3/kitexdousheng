package main

import (
	"context"
	"kitexdousheng/kitex_gen/feed"
)

// FeedSrvImpl implements the last service interface defined in the IDL.
type FeedSrvImpl struct{}

// Feed implements the FeedSrvImpl interface.
func (s *FeedSrvImpl) Feed(ctx context.Context, req *feed.DouyinFeedRequest) (resp *feed.DouyinFeedResponse, err error) {
	// TODO: Your code here...
	return
}
