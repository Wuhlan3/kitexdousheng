package service

import (
	"github.com/spf13/viper"
	"kitexdousheng/cmd/repository/db"
	"kitexdousheng/kitex_gen/feed"
	"kitexdousheng/kitex_gen/user"
	"time"
)

func PublishAction(uid int64, fileName string) error {
	b := []byte(fileName)
	imgName := string(b[:len(b)-3]) + "jpg"
	video := &db.Video{
		UId:            uid,
		PlayUrl:        fileName,
		CoverUrl:       imgName,
		CommentCount:   0,
		FavouriteCount: 0,
		Title:          fileName[:len(fileName)-4],
		CreateTime:     time.Now(),
		UpdateTime:     time.Now(),
		IsDeleted:      false,
	}
	if err := db.NewVideoDaoInstance().CreateVideo(video); err != nil {
		return err
	}

	return nil
}

func PublishList(uid int64) ([]*feed.Video, error) {
	videoPath := viper.GetString("cos.uriVideoPath")
	imgPath := viper.GetString("cos.uriPicturePath")
	videosList, err := db.NewVideoDaoInstance().QueryVideoListByUId(uid)
	var protoVideoList []*feed.Video
	for _, video := range *videosList {

		//follow获取
		tempFollow, err := db.NewFollowDaoInstance().QueryByUIdAndHisUId(uid, uid)
		var isFollow bool
		if err != nil {
			isFollow = false
			db.NewFollowDaoInstance().CreateFollow(&db.Follow{
				MyId:       uid,
				HisId:      uid,
				IsFollow:   false,
				CreateTime: time.Now(),
				UpdateTime: time.Now(),
			})
		} else {
			isFollow = tempFollow.IsFollow
		}
		//user获取
		tempUser, err := db.NewUserDaoInstance().QueryUserById(uid)
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
		isFavourite, err := db.NewFavouriteDaoInstance().QueryByVIdAndUId(video.Id, uid)
		if err != nil {
			isFavourite = false
			db.NewFavouriteDaoInstance().CreateFavourite(&db.Favourite{
				UId:         uid,
				VId:         video.Id,
				IsFavourite: false,
				CreateTime:  time.Now(),
				UpdateTime:  time.Now(),
			})
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
