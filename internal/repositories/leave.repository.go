package repositories

import (
	"context"

	"github.com/handarudwiki/payroll-sistem/internal/dto"
	"github.com/handarudwiki/payroll-sistem/internal/models"
	"github.com/handarudwiki/payroll-sistem/internal/utils"
	"gorm.io/gorm"
)

type (
	Leave interface {
		Create(ctx context.Context, leave models.Leave) (models.Leave, error)
		FindByID(ctx context.Context, id int) (models.Leave, error)
		Update(ctx context.Context, id int, leave models.Leave) (models.Leave, error)
		Delete(ctx context.Context, id int) error
		FindAll(ctx context.Context, base dto.BaseQuery) (leaves []models.Leave, totalData int64, err error)
	}

	leave struct {
		db *gorm.DB
	}
)

func NewLeaveRepository(db *gorm.DB) Leave {
	return &leave{
		db: db,
	}
}

func (r *leave) Create(ctx context.Context, leave models.Leave) (models.Leave, error) {
	err := r.db.Create(&leave).Error
	if err != nil {
		return models.Leave{}, err
	}
	return leave, nil
}

func (r *leave) FindByID(ctx context.Context, id int) (models.Leave, error) {
	var leave models.Leave
	err := r.db.Where("id = ?", id).First(&leave).Error
	if err != nil {
		return models.Leave{}, err
	}
	return leave, nil
}

func (r *leave) Update(ctx context.Context, id int, leave models.Leave) (models.Leave, error) {
	err := r.db.Model(&models.Leave{}).Where("id = ?", id).Updates(leave).Error
	if err != nil {
		return models.Leave{}, err
	}
	return leave, nil
}

func (r *leave) Delete(ctx context.Context, id int) error {
	err := r.db.Where("id = ?", id).Delete(&models.Leave{}).Error
	if err != nil {
		return err
	}
	return nil
}
func (r *leave) FindAll(ctx context.Context, base dto.BaseQuery) (leaves []models.Leave, totalData int64, err error) {
	err = r.db.Model(&models.Leave{}).Scopes(
		utils.Paginate(base.Page, base.Limit),
		utils.Search(base.Search),
	).Find(&leaves).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.Model(&models.Leave{}).Scopes(utils.Search(base.Search)).Count(&totalData).Error
	if err != nil {
		return nil, 0, err
	}

	return leaves, totalData, nil
}
