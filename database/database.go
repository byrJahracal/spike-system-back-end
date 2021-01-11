package database

import (
	"pku-class/market/data"
	eh "pku-class/market/error-handler"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	dsn := "root:123456@tcp(182.92.69.9:3306)/market?charset=utf8mb4&parseTime=True&loc=Local" //182.92.69.9
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	eh.ErrorHandler(err, "连接数据库失败！", "数据库连接成功！")
	DB.AutoMigrate(&data.User{})
	DB.AutoMigrate(&data.Commodity{})
	DB.AutoMigrate(&data.Order{})
	DB.AutoMigrate(&data.Menu{})
}
