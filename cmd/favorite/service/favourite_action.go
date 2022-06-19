package service

import (
	"errors"
	"kitexdousheng/cmd/repository/db"
	"time"
)

type FavouriteActionFlow struct {
	UId        int64
	VId        int64
	ActionType int32
}

func FavouriteAction(userId int64, videoId int64, actionType int32) error {
	return NewFavouriteActionFlow(userId, videoId, actionType).Do()
}

func NewFavouriteActionFlow(userId int64, videoId int64, actionType int32) *FavouriteActionFlow {
	return &FavouriteActionFlow{
		UId:        userId,
		VId:        videoId,
		ActionType: actionType,
	}
}

func (f *FavouriteActionFlow) Do() error {
	if err := f.checkParam(); err != nil {
		return err
	}
	err := f.action()
	if err != nil {
		return err
	}
	return nil
}

func (f *FavouriteActionFlow) checkParam() error {
	return nil
}

func (f *FavouriteActionFlow) action() error {
	isFavourite, err := db.NewFavouriteDaoInstance().QueryByVIdAndUId(f.VId, f.UId)
	//这里应该对错误处理再细分一下，分成未知错误和查不到，后面再改
	if err != nil {
		//没有找到
		//取消点赞
		if f.ActionType != 1 {
			return nil
		}
		//后面可以改成事务加goroutine解决
		if err := db.NewFavouriteDaoInstance().CreateFavourite(&db.Favourite{
			UId:         f.UId,
			VId:         f.VId,
			IsFavourite: true,
			CreateTime:  time.Now(),
			UpdateTime:  time.Now(),
		}); err != nil {
			return err
		}

		err := db.NewVideoDaoInstance().IncFavouriteCount(f.VId)
		if err != nil {
			return err
		}

		return nil
	}
	//找到的情况，记得区分还要不要改
	if f.ActionType == 1 && isFavourite {
		return nil
	}
	if f.ActionType == 2 && !isFavourite {
		return nil
	}

	if isFavourite == true {
		err := db.NewVideoDaoInstance().DecFavouriteCount(f.VId)
		if err != nil {
			return err
		}
	} else {
		err := db.NewVideoDaoInstance().IncFavouriteCount(f.VId)
		if err != nil {
			return err
		}
	}
	err = db.NewFavouriteDaoInstance().UpdateIsFavourite(f.VId, f.UId, !isFavourite)

	if err != nil {
		return errors.New("修改失败")
	}
	return nil
}
