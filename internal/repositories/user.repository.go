package repositories

import (
	"github.com/handarudwiki/payroll-sistem/internal/models"
	"gorm.io/gorm"
)

type (
	User interface {
		FindByID(id int) (models.User, error)
		FindByUsername(username string) (models.User, error)
		Create(user models.User) (models.User, error)
	}

	userRepository struct {
		db *gorm.DB
	}
)

func NewUserRepository(db *gorm.DB) User {
	return &userRepository{
		db: db,
	}
}
func (r *userRepository) FindByID(id int) (models.User, error) {
	var user models.User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *userRepository) FindByUsername(username string) (models.User, error) {
	var user models.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *userRepository) Create(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
