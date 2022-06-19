package service

import (
	"kitexdousheng/cmd/repository/db"
	"kitexdousheng/kitex_gen/feed"
	"kitexdousheng/kitex_gen/user"
	"time"
)

func PublishAction(uid int64, fileName string) error {
	len := len(fileName)
	video := &db.Video{
		UId:            uid,
		PlayUrl:        fileName,
		CoverUrl:       "bear.jpg",
		CommentCount:   10,
		FavouriteCount: 10,
		Title:          fileName[:len-4],
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
	videosList, err := db.NewVideoDaoInstance().QueryVideoListByUId(uid)
	var protoVideoList []*feed.Video
	for _, video := range *videosList {
		tempUser, err := db.NewUserDaoInstance().QueryUserById(uid)
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
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavouriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    false,
			Title:         video.Title,
		})
	}

	return protoVideoList, err
}
