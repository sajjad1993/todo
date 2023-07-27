package orm

import "gorm.io/gorm"

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&TodoList{},
		&TodoItem{},
	)
	if err != nil {
		return err
	}
	return nil
}
