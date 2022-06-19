package service

import "kitexdousheng/kitex_gen/user"
import "kitexdousheng/cmd/repository/db"

func RelationFollowList(uid int64) ([]*user.User, error) {
	var protoUserList []*user.User
	followList, err := db.NewFollowDaoInstance().QueryByUId(uid)
	if err != nil {
		return nil, err
	}
	for _, follow := range *followList {
		User, err := db.NewUserDaoInstance().QueryUserById(follow.HisId)
		if err != nil {
			return nil, err
		}
		protoUserList = append(protoUserList, &user.User{
			Id:            User.Id,
			Name:          User.Name,
			FollowCount:   &User.FollowCount,
			FollowerCount: &User.FollowerCount,
			IsFollow:      true,
		})
	}
	return protoUserList, nil
}
