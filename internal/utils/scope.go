package utils

import (
	"time"

	"gorm.io/gorm"
)

func Search(search string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if search == "" {
			return db
		}
		return db.Where("name ILIKE ?", "%"+search+"%")
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

func FilterEmployeeID(employeeID *int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if employeeID == nil {
			return db
		}

		return db.Where("employee_id = ?", *employeeID)
	}
}

func FilterPeriodMonth(period *time.Time) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if period == nil {
			return db
		}

		return db.Where("EXTRACT(MONTH FROM period) = ? AND EXTRACT(YEAR FROM period) = ?",
			period.Month(), period.Year())
	}
}
