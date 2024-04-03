package testutil

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

func InitDB() (*gorm.DB, error) {
	// TEST_DB="user:password@tcp(localhost:3306)/database?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := os.Getenv("TEST_DB")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: gormlogger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			gormlogger.Config{
				SlowThreshold:             time.Millisecond * 100,
				IgnoreRecordNotFoundError: true,
				LogLevel:                  gormlogger.Info,
			},
		),
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}
