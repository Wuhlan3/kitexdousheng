package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"kitexdousheng/pkg/constants"
)

var DB *gorm.DB

// Init init DB
func Init() {
	var err error
	DB, err = gorm.Open(mysql.Open(constants.MySQLDefaultDSN),
		&gorm.Config{
			PrepareStmt:            true, //预编译
			SkipDefaultTransaction: true, //跳过默认事务
		},
	)
	if err != nil {
		panic(err)
	}

	//if err = DB.Use(gormopentracing.New()); err != nil {
	//	panic(err)
	//}

	m := DB.Migrator()
	if m.HasTable(&User{}) {
		return
	}
	if err = m.CreateTable(&User{}); err != nil {
		panic(err)
	}
}
