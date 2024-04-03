package gormutil

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gorm.io/gorm"

	"github.com/liangjunmo/go-code-snippets/internal/testutil"
)

type User struct {
	gorm.Model
	Username string
}

func (*User) TableName() string {
	return "user"
}

type User2 struct {
	gorm.Model
	Username string
	Password string
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

	err = RecreateTables(db, []interface{}{&User2{}})
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
