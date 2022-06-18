package db

import (
	"fmt"
	"gorm.io/gorm"
	"log"
	"sync"
	"time"
)

type User struct {
	Id            int64     `gorm:"column:id"`
	Name          string    `gorm:"column:name"`
	Password      string    `gorm:"column:password"`
	FollowCount   int64     `gorm:"column:follow_count"`
	FollowerCount int64     `gorm:"column:follower_count"`
	CreateTime    time.Time `gorm:"column:create_time"`
	UpdateTime    time.Time `gorm:"column:update_time"`
	IsDeleted     bool      `gorm:"column:is_deleted"`
}

func (User) TableName() string {
	return "user"
}

type UserDao struct {
}

var userDao *UserDao //DAO(DataAccessObject)模式
var userOnce sync.Once

//var db = DB

func NewUserDaoInstance() *UserDao {
	userOnce.Do(
		func() {
			userDao = &UserDao{}
		})
	return userDao
}

func (*UserDao) CreateUser(user *User) error {

	if err := DB.Create(user).Error; err != nil {
		log.Println("insert user err:" + err.Error())
		return err
	}
	return nil
}

func (*UserDao) QueryUserById(id int64) (*User, error) {
	var user User
	err := DB.Where("id = ?", id).Find(&user).Error
	if err != nil {
		log.Println("find user by id err:" + err.Error())
		return nil, err
	}
	return &user, nil

}

func (*UserDao) QueryUserByName(name string) (*User, error) {
	var user User
	//fmt.Println(name)
	err := DB.Where("name = ?", name).Find(&user).Error
	//fmt.Println("hello")
	if err != nil {
		fmt.Println("hello")
		log.Println("find user by name err:" + err.Error())
		return nil, err
	}
	return &user, nil

}

func (*UserDao) IncUserFollow(uid int64) error {
	err := DB.Model(User{}).Where("id = ?", uid).UpdateColumn("follower_count", gorm.Expr("follow_count + ?", 1)).Error
	if err != nil {
		log.Println("inc user follow count error")
		return err
	}
	return nil
}

func (*UserDao) DecUserFollow(uid int64) error {
	err := DB.Model(User{}).Where("id = ?", uid).UpdateColumn("follower_count", gorm.Expr("follow_count - ?", 1)).Error
	if err != nil {
		log.Println("dec user follow count error")
		return err
	}
	return nil
}

func (*UserDao) IncUserFollower(uid int64) error {
	err := DB.Model(User{}).Where("id = ?", uid).UpdateColumn("follower_count", gorm.Expr("follower_count + ?", 1)).Error
	if err != nil {
		log.Println("inc user follower count error")
		return err
	}
	return nil
}

func (*UserDao) DecUserFollower(uid int64) error {
	err := DB.Model(User{}).Where("id = ?", uid).UpdateColumn("follower_count", gorm.Expr("follower_count - ?", 1)).Error
	if err != nil {
		log.Println("dec user follower count error")
		return err
	}
	return nil
}
