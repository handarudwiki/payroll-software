package utils_test

import (
	"github.com/handarudwiki/payroll-sistem/internal/models"
	"gorm.io/gorm"
)

func CreateUserTest(db *gorm.DB, name, username, password string, role models.UserRole) (models.User, error) {

	user := models.User{
		Name:     name,
		Username: username,
		Password: password,
		Role:     role,
	}

	user.EncryptPassword()

	if err := db.Create(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}
