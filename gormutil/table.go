package gormutil

import (
	"fmt"

	"gorm.io/gorm"
)

func RecreateTables(db *gorm.DB, models []interface{}) error {
	stmt := &gorm.Statement{DB: db}
	for _, model := range models {
		err := stmt.Parse(model)
		if err != nil {
			return err
		}

		err = stmt.Migrator().DropTable(model)
		if err != nil {
			return err
		}

		err = stmt.AutoMigrate(model)
		if err != nil {
			return err
		}

		if !stmt.Migrator().HasTable(model) {
			return fmt.Errorf("failed to table create: %#v", stmt.Table)
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

		err = stmt.AutoMigrate(model)
		if err != nil {
			return err
		}

		sql := fmt.Sprintf("truncate table %s;", stmt.Table)
		err = stmt.Exec(sql).Error
		if err != nil {
			return err
		}
	}
	return nil
}
