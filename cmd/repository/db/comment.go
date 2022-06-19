package db

import (
	"log"
	"sync"
	"time"
)

type Comment struct {
	Id         int64     `gorm:"column:id"`
	VId        int64     `gorm:"column:vid"`
	UId        int64     `gorm:"column:uid"`
	Content    string    `gorm:"column:content"`
	CreateTime time.Time `gorm:"column:create_time"`
	UpdateTime time.Time `gorm:"column:update_time"`
	IsDeleted  bool      `gorm:"column:is_deleted"`
}

func (Comment) TableName() string {
	return "comment"
}

type CommentDao struct {
}

var commentDao *CommentDao //DAO(DataAccessObject)模式
var commentOnce sync.Once

func NewCommentDaoInstance() *CommentDao {
	commentOnce.Do(
		func() {
			commentDao = &CommentDao{}
		})
	return commentDao
}

func (c *CommentDao) QueryByVId(vid int64) (*[]Comment, error) {
	var comments []Comment
	err := DB.Where("vid = ?", vid).Find(&comments).Error
	if err != nil {
		log.Println("find comment by vid err:" + err.Error())
		return nil, err
	}
	return &comments, nil
}

func (c *CommentDao) CreateComment(content *Comment) error {
	if err := DB.Create(content).Error; err != nil {
		log.Println("insert favourite err:" + err.Error())
		return err
	}
	return nil
}
