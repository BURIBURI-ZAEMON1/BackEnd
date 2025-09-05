package database

import (
	"backend/conf/config"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	host := config.Config.GetString("mysql.host")
	port := config.Config.GetString("mysql.port")
	user := config.Config.GetString("mysql.user")
	password := config.Config.GetString("mysql.password")
	DBname := config.Config.GetString("mysql.DBname")

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
		user, password, host, port, DBname)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	err = autoMigrate(db)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("数据库连接成功")

	DB = db
}
