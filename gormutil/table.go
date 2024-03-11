package gormutil

import (
	"fmt"

	"gorm.io/gorm"
)

func DropCreateTables(db *gorm.DB, models []interface{}) error {
	for _, model := range models {
		err := db.Migrator().DropTable(model)
		if err != nil {
			return err
		}

		err = db.AutoMigrate(model)
		if err != nil {
			return err
		}

		if !db.Migrator().HasTable(model) {
			return fmt.Errorf("table not created: %#v", model)
		}
	}
	return nil
}

func TruncateTables(db *gorm.DB, models []interface{}) error {
	stmt := &gorm.Statement{DB: db}
	for _, model := range models {
		err := stmt.Parse(model)
		if err != nil {
			return err
		}

		err = db.AutoMigrate(model)
		if err != nil {
			return err
		}

		sql := fmt.Sprintf("truncate table %s;", stmt.Schema.Table)
		err = db.Exec(sql).Error
		if err != nil {
			return err
		}
	}
	return nil
}
