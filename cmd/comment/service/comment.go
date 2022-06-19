package service

import (
	"kitexdousheng/cmd/repository/db"
	"time"
)

type CommentActionFlow struct {
	UId     int64
	VId     int64
	Content string
}

func CommentAction(userId int64, videoId int64, content string) error {
	return NewCommentActionFlow(userId, videoId, content).Do()
}

func NewCommentActionFlow(userId int64, videoId int64, content string) *CommentActionFlow {
	return &CommentActionFlow{
		UId:     userId,
		VId:     videoId,
		Content: content,
	}
}

func (c *CommentActionFlow) Do() error {
	if err := c.checkParam(); err != nil {
		return err
	}
	err := c.action()
	if err != nil {
		return err
	}
	return nil
}

func (c *CommentActionFlow) checkParam() error {
	return nil
}

func (c *CommentActionFlow) action() error {
	err := db.NewCommentDaoInstance().CreateComment(&db.Comment{
		VId:        c.VId,
		UId:        c.UId,
		Content:    c.Content,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		IsDeleted:  false,
	})
	if err != nil {
		return err
	}
	//增加评论数
	err = db.NewVideoDaoInstance().IncCommentCount(c.VId)
	if err != nil {
		return err
	}
	return nil
}
