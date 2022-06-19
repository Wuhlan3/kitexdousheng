package service

import (
	"github.com/spf13/viper"
	"kitexdousheng/cmd/repository/db"
	"kitexdousheng/kitex_gen/feed"
	"kitexdousheng/kitex_gen/user"
	"strconv"
)

//Feed 没有建立起和myUId之间的关系，后续优化可以加入推荐算法
func Feed(myUId int64) ([]*feed.Video, error) {
	videoPath := viper.GetString("cos.uriVideoPath")
	imgPath := viper.GetString("cos.uriPicturePath")
	maxNumStr := viper.GetString("video.maxNumPerTimes")
	maxNum, err := strconv.ParseInt(maxNumStr, 10, 64)
	if err != nil {
		return nil, err
	}
	videosList, err := db.NewVideoDaoInstance().QueryVideoList(int(maxNum))
	var protoVideoList []*feed.Video
	for _, video := range *videosList {
		tempUser, err := db.NewUserDaoInstance().QueryUserById(video.UId)
		if err != nil {
			return nil, err
		}
		tempProtoUser := &user.User{
			Id:            tempUser.Id,
			Name:          tempUser.Name,
			FollowCount:   &tempUser.FollowCount,
			FollowerCount: &tempUser.FollowerCount,
			IsFollow:      false,
		}
		protoVideoList = append(protoVideoList, &feed.Video{
			Id:            video.Id,
			Author:        tempProtoUser,
			PlayUrl:       videoPath + video.PlayUrl,
			CoverUrl:      imgPath + video.CoverUrl,
			FavoriteCount: video.FavouriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    false,
			Title:         video.Title,
		})
	}
	return protoVideoList, err
}
