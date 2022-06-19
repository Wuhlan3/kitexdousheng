package service

import (
	"fmt"
	"github.com/spf13/viper"
	"kitexdousheng/cmd/repository/db"
	"kitexdousheng/kitex_gen/feed"
	"kitexdousheng/kitex_gen/user"
	"time"
)

func FavoriteList(uid int64) ([]*feed.Video, error) {
	path := viper.GetString("video.absolutePath")
	favouriteList, err := db.NewFavouriteDaoInstance().QueryByUId(uid)
	var protoVideoList []*feed.Video
	for _, fav := range *favouriteList {
		if fav.IsFavourite {
			User, err := db.NewUserDaoInstance().QueryUserById(uid)
			if err != nil {
				return nil, err
			}

			video, err := db.NewVideoDaoInstance().QueryVideoById(fav.VId)
			if err != nil {
				return nil, err
			}

			var IsFollow bool

			follow, err := db.NewFollowDaoInstance().QueryByUIdAndHisUId(uid, video.UId)
			if err != nil {
				db.NewFollowDaoInstance().CreateFollow(&db.Follow{
					MyId:       uid,
					HisId:      video.Id,
					IsFollow:   false,
					CreateTime: time.Now(),
					UpdateTime: time.Now(),
				})
				IsFollow = false
			} else {
				IsFollow = follow.IsFollow
			}
			demoUser := &user.User{
				Id:            User.Id,
				Name:          User.Name,
				FollowCount:   &User.FollowCount,
				FollowerCount: &User.FollowerCount,
				IsFollow:      IsFollow,
			}

			fmt.Println(path + ":" + video.PlayUrl)
			protoVideoList = append(protoVideoList, &feed.Video{
				Id:            video.Id,
				Author:        demoUser,
				PlayUrl:       path + video.PlayUrl,
				CoverUrl:      path + video.CoverUrl,
				FavoriteCount: video.FavouriteCount,
				CommentCount:  video.CommentCount,
				IsFavorite:    true,
				Title:         video.Title,
			})
		}
	}

	return protoVideoList, err
}
