package service

import (
	"fmt"
	"kitexdousheng/cmd/repository/db"
	"kitexdousheng/kitex_gen/feed"
	"kitexdousheng/kitex_gen/user"
	"log"
	"strconv"
	"time"

	"github.com/spf13/viper"
)

//Feed 后续优化可以加入推荐算法
func Feed(myUId int64, latestTime int64) ([]*feed.Video, error) {
	videoPath := viper.GetString("cos.uriVideoPath")
	imgPath := viper.GetString("cos.uriPicturePath")
	maxNumStr := viper.GetString("video.maxNumPerTimes")
	maxNum, err := strconv.ParseInt(maxNumStr, 10, 64)
	if err != nil {
		return nil, err
	}

	//查询Redis中是否存在 视频序号序列ZSet
	VIdList, err := GetVIdListFromRedis(latestTime, maxNum)
	if err != nil {
		return nil, err
	}

	//如果存在，根据序号来获取视频相关信息，Hash
	videosList, err := GetVideoListFromRedis(VIdList)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	var protoVideoList []*feed.Video
	for _, video := range videosList {
		//follow获取
		tempFollow, err := db.NewFollowDaoInstance().QueryByUIdAndHisUId(myUId, video.UId)
		var isFollow bool
		if err != nil {
			isFollow = false
			if err := db.NewFollowDaoInstance().CreateFollow(&db.Follow{
				MyId:       myUId,
				HisId:      video.UId,
				IsFollow:   false,
				CreateTime: time.Now(),
				UpdateTime: time.Now(),
			}); err != nil {
				log.Println(err.Error())
			}
		} else {
			isFollow = tempFollow.IsFollow
		}

		//user获取
		tempUser, err := db.NewUserDaoInstance().QueryUserById(video.UId)
		if err != nil {
			return nil, err
		}
		tempProtoUser := &user.User{
			Id:            tempUser.Id,
			Name:          tempUser.Name,
			FollowCount:   &tempUser.FollowCount,
			FollowerCount: &tempUser.FollowerCount,
			IsFollow:      isFollow,
		}

		//favourite获取
		isFavourite, err := db.NewFavouriteDaoInstance().QueryByVIdAndUId(video.Id, myUId)
		if err != nil {
			isFavourite = false
			if err := db.NewFavouriteDaoInstance().CreateFavourite(&db.Favourite{
				UId:         myUId,
				VId:         video.Id,
				IsFavourite: false,
				CreateTime:  time.Now(),
				UpdateTime:  time.Now(),
			}); err != nil {
				log.Println(err.Error())
			}
		}
		//生成videoList
		protoVideoList = append(protoVideoList, &feed.Video{
			Id:            video.Id,
			Author:        tempProtoUser,
			PlayUrl:       videoPath + video.PlayUrl,
			CoverUrl:      imgPath + video.CoverUrl,
			FavoriteCount: video.FavouriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    isFavourite,
			Title:         video.Title,
		})
	}
	return protoVideoList, err
}
