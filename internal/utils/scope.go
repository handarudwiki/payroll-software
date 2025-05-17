package utils

import "gorm.io/gorm"

func Search(search string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if search == "" {
			return db
		}
		return db.Where("name LIKE ?", "%"+search+"%")
	}
}
