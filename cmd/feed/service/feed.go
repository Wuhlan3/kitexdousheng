package service

import (
	"fmt"
	"github.com/spf13/viper"
	"kitexdousheng/cmd/feed/repository/db"
	"kitexdousheng/kitex_gen/feed"
	"kitexdousheng/kitex_gen/user"
	"strconv"
)

func Feed(myUId int64) ([]*feed.Video, error) {
	path := viper.GetString("cos.uriPath")
	maxNumStr := viper.GetString("video.maxNumPerTimes")
	maxNum, err := strconv.ParseInt(maxNumStr, 10, 64)
	if err != nil {
		return nil, err
	}
	videosList, err := db.NewVideoDaoInstance().QueryVideoList(int(maxNum))
	var protoVideoList []*feed.Video
	for _, video := range *videosList {
		fmt.Println(path + video.PlayUrl)
		protoVideoList = append(protoVideoList, &feed.Video{
			Id:            video.Id,
			Author:        &user.User{},
			PlayUrl:       path + video.PlayUrl,
			CoverUrl:      path + video.CoverUrl,
			FavoriteCount: video.FavouriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    false,
			Title:         video.Title,
		})
	}

	return protoVideoList, err
}
