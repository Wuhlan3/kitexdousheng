package main

import (
	"context"
	"kitexdousheng/cmd/favorite/service"
	"kitexdousheng/kitex_gen/favorite"
	"kitexdousheng/pkg/errno"
	"log"
)

// FavoriteSrvImpl implements the last service interface defined in the IDL.
type FavoriteSrvImpl struct{}

// FavoriteAction implements the FavoriteSrvImpl interface.
func (s *FavoriteSrvImpl) FavoriteAction(ctx context.Context, req *favorite.DouyinFavoriteActionRequest) (*favorite.DouyinFavoriteActionResponse, error) {
	log.Println(req)
	err := service.FavouriteAction(req.UserId, req.VideoId, req.ActionType)
	if err != nil {
		return &favorite.DouyinFavoriteActionResponse{
			StatusCode: errno.ServiceErr.ErrCode,
			StatusMsg:  &errno.ServiceErr.ErrMsg,
		}, err
	}
	return &favorite.DouyinFavoriteActionResponse{
		StatusCode: 0,
		StatusMsg:  &errno.Success.ErrMsg,
	}, err
}

// FavoriteList implements the FavoriteSrvImpl interface.
func (s *FavoriteSrvImpl) FavoriteList(ctx context.Context, req *favorite.DouyinFavoriteListRequest) (*favorite.DouyinFavoriteListResponse, error) {
	videos, err := service.FavoriteList(req.UserId)
	if err != nil {
		return &favorite.DouyinFavoriteListResponse{
			StatusCode: errno.ServiceErr.ErrCode,
			StatusMsg:  &errno.ServiceErr.ErrMsg,
		}, err
	}
	return &favorite.DouyinFavoriteListResponse{
		StatusCode: 0,
		StatusMsg:  &errno.Success.ErrMsg,
		VideoList:  videos,
	}, err
}
