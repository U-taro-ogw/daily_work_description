package connect

import (
	"os"

	"github.com/jinzhu/gorm"

	"github.com/U-taro-ogw/daily_work_description/work_api/models"
)

func MysqlConnection() *gorm.DB {
	USER := os.Getenv("MYSQL_USER")
	PASS := os.Getenv("MYSQL_ROOT_PASSWORD")
	PROTOCOL := "tcp(auth_db:3306)"
	DBNAME := os.Getenv("MYSQL_DATABASE")
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
	db, err := gorm.Open("mysql", CONNECT)

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.WorkRecord{})

	return db
}
