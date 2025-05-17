package repositories

import (
	"context"

	"github.com/handarudwiki/payroll-sistem/internal/models"
	"gorm.io/gorm"
)

type (
	User interface {
		FindByID(ctx context.Context, id int) (models.User, error)
		FindByUsername(ctx context.Context, username string) (models.User, error)
		Create(ctx context.Context, user models.User) (models.User, error)
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
func (r *userRepository) FindByID(ctx context.Context, id int) (models.User, error) {
	var user models.User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *userRepository) FindByUsername(ctx context.Context, username string) (models.User, error) {
	var user models.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *userRepository) Create(ctx context.Context, user models.User) (models.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
