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

func FilterDepartment(departmentID *int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if departmentID == nil {
			return db
		}
		return db.Where("department_id = ?", *departmentID)
	}
}

func FilterPosition(positionID *int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if positionID == nil {
			return db
		}
		return db.Where("position_id = ?", *positionID)
	}
}
