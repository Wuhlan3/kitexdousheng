package service

import (
	"kitexdousheng/cmd/repository/db"
	"time"
)

type RelationActionFlow struct {
	MyUId      int64
	HisUId     int64
	ActionType string
}

func RelationAction(myUId int64, hisUId int64, actionType string) error {
	return NewRelationActionFlow(myUId, hisUId, actionType).Do()
}

func NewRelationActionFlow(myUId int64, hisUId int64, actionType string) *RelationActionFlow {
	return &RelationActionFlow{
		MyUId:      myUId,
		HisUId:     hisUId,
		ActionType: actionType,
	}
}

func (c *RelationActionFlow) Do() error {
	if err := c.checkParam(); err != nil {
		return err
	}
	err := c.action()
	if err != nil {
		return err
	}
	return nil
}

func (c *RelationActionFlow) checkParam() error {
	return nil
}

func (c *RelationActionFlow) action() error {
	var isFollow bool
	if c.ActionType == "1" {
		isFollow = true
	} else {
		isFollow = false
	}
	follow, err := db.NewFollowDaoInstance().QueryByUIdAndHisUId(c.MyUId, c.HisUId)
	if err != nil || follow == nil {
		err := db.NewFollowDaoInstance().CreateFollow(&db.Follow{
			MyId:       c.MyUId,
			HisId:      c.HisUId,
			IsFollow:   isFollow,
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
		})
		if err != nil {
			return err
		}
		return nil
	}
	if follow.IsFollow == true && isFollow == false {
		err := db.NewUserDaoInstance().DecUserFollow(c.MyUId)
		if err != nil {
			return err
		}
		err = db.NewUserDaoInstance().DecUserFollower(c.HisUId)
		if err != nil {
			return err
		}
	} else if follow.IsFollow == false && isFollow == true {
		err := db.NewUserDaoInstance().IncUserFollow(c.MyUId)
		if err != nil {
			return err
		}
		err = db.NewUserDaoInstance().IncUserFollower(c.HisUId)
		if err != nil {
			return err
		}
	}

	err = db.NewFollowDaoInstance().UpdateFollow(c.MyUId, c.HisUId, isFollow)
	if err != nil {
		return err
	}
	return nil
}
