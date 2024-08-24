package driver

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"starfruit.top/user_srv/model"
)

var DB *gorm.DB

func init() {
	dsn := "root:12345678@tcp(127.0.0.1:3306)/starfruit?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		panic(err)
	}

	DB = db
}
