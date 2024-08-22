package examples

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type Product struct {
	gorm.Model
	Code string `gorm:"unique" json:"code"`
	Name string `gorm:"unique" json:"name"`
}

func main() {
	dns := "root:12345678@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=True&loc=Local"
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Silent,
			Colorful:      true,
		},
	)

	var db, err = gorm.Open(mysql.Open(dns), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	_ = db.AutoMigrate(&Product{})

	db.Create(&Product{Code: "123", Name: "tom"})

	var product Product
	db.First(&product, 1)

	db.Model(&product).Update("code", "456")

	fmt.Println(product.Model.CreatedAt)

}
