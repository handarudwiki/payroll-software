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
		Update(ctx context.Context, user models.User, id int) (models.User, error)
		UpdatePassword(ctx context.Context, id int, password string) (models.User, error)
		BulkCreate(ctx context.Context, users []models.User) ([]models.User, error)
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

func (r *userRepository) Update(ctx context.Context, user models.User, id int) (models.User, error) {
	err := r.db.Model(&models.User{}).Where("id = ?", id).Updates(user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
func (r *userRepository) UpdatePassword(ctx context.Context, id int, password string) (models.User, error) {
	var user models.User
	err := r.db.Model(&models.User{}).Where("id = ?", id).Update("password", password).First(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *userRepository) BulkCreate(ctx context.Context, users []models.User) ([]models.User, error) {
	if len(users) == 0 {
		return nil, nil
	}

	err := r.db.Create(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
