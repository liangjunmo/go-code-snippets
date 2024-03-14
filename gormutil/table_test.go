package gormutil

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"gorm.io/gorm"

	"github.com/liangjunmo/goutil/internal/testutil"
)

type User struct {
	UID        uint32         `gorm:"column:id;type:int unsigned;not null;auto_increment;primary key;" json:"-"`
	CreateTime time.Time      `gorm:"column:create_time;type:datetime;not null;autoCreateTime;" json:"-"`
	UpdateTime time.Time      `gorm:"column:update_time;type:datetime;not null;autoUpdateTime;" json:"-"`
	DeleteTime gorm.DeletedAt `gorm:"column:delete_time;type:datetime;default:null;index:idx_delete_time;" json:"-"`
	Username   string         `gorm:"column:username;type:varchar(32);not null;index:idx_username,unique;" json:"-"`
}

func (*User) TableName() string {
	return "user"
}

type User2 struct {
	UID        uint32         `gorm:"column:id;type:int unsigned;not null;auto_increment;primary key;" json:"-"`
	CreateTime time.Time      `gorm:"column:create_time;type:datetime;not null;autoCreateTime;" json:"-"`
	UpdateTime time.Time      `gorm:"column:update_time;type:datetime;not null;autoUpdateTime;" json:"-"`
	DeleteTime gorm.DeletedAt `gorm:"column:delete_time;type:datetime;default:null;index:idx_delete_time;" json:"-"`
	Username   string         `gorm:"column:username;type:varchar(32);not null;index:idx_username,unique;" json:"-"`
	Password   string         `gorm:"column:password;type:varchar(100);not null;" json:"-"`
}

func (*User2) TableName() string {
	return "user"
}

func TestDropCreateTables(t *testing.T) {
	db, err := testutil.InitDB()
	require.Nil(t, err)
	defer func() {
		db, _ := db.DB()
		db.Close()
	}()

	err = db.Migrator().DropTable(&User{})
	require.Nil(t, err)

	err = db.AutoMigrate(&User{})
	require.Nil(t, err)

	err = DropCreateTables(db, []interface{}{&User2{}})
	require.Nil(t, err)

	exist := db.Migrator().HasColumn(&User2{}, "password")
	require.True(t, exist)
}

func TestTruncateTables(t *testing.T) {
	db, err := testutil.InitDB()
	require.Nil(t, err)
	defer func() {
		db, _ := db.DB()
		db.Close()
	}()

	err = db.Migrator().DropTable(&User{})
	require.Nil(t, err)

	err = db.AutoMigrate(&User{})
	require.Nil(t, err)

	err = db.Create(&User{Username: "test"}).Error
	require.Nil(t, err)

	err = TruncateTables(db, []interface{}{&User{}})
	require.Nil(t, err)

	var count int64
	err = db.Model(&User{}).Count(&count).Error
	require.Nil(t, err)
	require.True(t, count == 0)
}
