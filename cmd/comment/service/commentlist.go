package service

import (
	"kitexdousheng/cmd/repository/db"
	"kitexdousheng/kitex_gen/comment"
	"kitexdousheng/kitex_gen/user"
)

func CommentList(vid int64) ([]*comment.Comment, error) {

	commentList, err := db.NewCommentDaoInstance().QueryByVId(vid)
	if err != nil {
		return nil, err
	}
	var protoCommentList []*comment.Comment
	for _, tempComment := range *commentList {
		tempUser, err := db.NewUserDaoInstance().QueryUserById(tempComment.UId)
		if err != nil {
			return nil, err
		}
		demoUser := &user.User{
			Id:            tempUser.Id,
			Name:          tempUser.Name,
			FollowCount:   &tempUser.FollowCount,
			FollowerCount: &tempUser.FollowerCount,
			IsFollow:      false,
		}
		month := tempComment.CreateTime.Format("01")
		date := tempComment.CreateTime.Format("02")

		protoCommentList = append(protoCommentList, &comment.Comment{
			Id:         tempComment.Id,
			User:       demoUser,
			Content:    tempComment.Content,
			CreateDate: month + "-" + date,
		})
	}
	return protoCommentList, nil
}
