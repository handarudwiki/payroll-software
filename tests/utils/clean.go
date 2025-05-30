package utils_test

import "gorm.io/gorm"

func CleanAllTables(db *gorm.DB) error {

	err := db.Exec("DELETE FROM users").Error

	if err != nil {
		return err
	}

	return nil
}
