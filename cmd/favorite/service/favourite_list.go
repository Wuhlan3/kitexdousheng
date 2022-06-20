package service

import (
	"github.com/spf13/viper"
	"kitexdousheng/cmd/repository/db"
	"kitexdousheng/kitex_gen/feed"
	"kitexdousheng/kitex_gen/user"
	"time"
)

func FavoriteList(uid int64) ([]*feed.Video, error) {
	videoPath := viper.GetString("cos.uriVideoPath")
	imgPath := viper.GetString("cos.uriPicturePath")
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

			protoVideoList = append(protoVideoList, &feed.Video{
				Id:            video.Id,
				Author:        demoUser,
				PlayUrl:       videoPath + video.PlayUrl,
				CoverUrl:      imgPath + video.CoverUrl,
				FavoriteCount: video.FavouriteCount,
				CommentCount:  video.CommentCount,
				IsFavorite:    true,
				Title:         video.Title,
			})
		}
	}

	return protoVideoList, err
}
