package gormutil

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"

	"github.com/liangjunmo/go-code-snippets/trace"
)

func Connect(dsn string, config *gorm.Config) (*gorm.DB, error) {
	if config == nil {
		config = &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
			Logger: trace.NewGORMLogger(
				gormlogger.Config{
					SlowThreshold:             time.Millisecond * 100,
					IgnoreRecordNotFoundError: true,
					LogLevel:                  gormlogger.Info,
				},
			),
		}
	}
	return gorm.Open(mysql.Open(dsn), config)
}
