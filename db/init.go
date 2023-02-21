package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

const (
	dbuser = "root"
	pwd    = "123456"
	host   = "127.0.0.1"
	port   = 3306
	dbname = "db_course"
)

var DB *gorm.DB

func Init() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbuser, pwd, host, port, dbname,
	)
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt: true,
			Logger:      logger.Default.LogMode(logger.Info),
		},
	)
	if err != nil {
		panic("DB init failed: " + err.Error())
	}

	// 数据自动迁移
	if err = DB.AutoMigrate(&Student{}); err != nil {
		panic("Data Auto Migrate failed :" + err.Error())
	}

}

func ParseTime(s string) (*time.Time, error) {
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return nil, err
	}
	return &t, nil
}
