package service

import "kitexdousheng/kitex_gen/user"
import "kitexdousheng/cmd/repository/db"

func RelationFollowerList(uid int64) ([]*user.User, error) {
	var protoUserList []*user.User
	followerList, err := db.NewFollowDaoInstance().QueryByHisUId(uid)
	if err != nil {
		return nil, err
	}
	for _, follow := range *followerList {
		User, err := db.NewUserDaoInstance().QueryUserById(follow.MyId)
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
