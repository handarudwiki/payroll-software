package repositories

import (
	"context"

	"github.com/handarudwiki/payroll-sistem/internal/dto"
	"github.com/handarudwiki/payroll-sistem/internal/models"
	"github.com/handarudwiki/payroll-sistem/internal/utils"
	"gorm.io/gorm"
)

type (
	Position interface {
		Create(ctx context.Context, position models.Position) (models.Position, error)
		FindByID(ctx context.Context, id int) (models.Position, error)
		Update(ctx context.Context, id int, position models.Position) (models.Position, error)
		Delete(ctx context.Context, id int) error
		FindAll(ctx context.Context, base dto.BaseQuery) (positions []models.Position, totalData int64, err error)
		BulkCreate(ctx context.Context, positions []models.Position) ([]models.Position, error)
		FindAllOnly(ctx context.Context) (positions []models.Position, err error)
	}

	position struct {
		db *gorm.DB
	}
)

func NewPositionRepository(db *gorm.DB) Position {
	return &position{
		db: db,
	}
}

func (r *position) Create(ctx context.Context, position models.Position) (models.Position, error) {
	err := r.db.Create(&position).Error
	if err != nil {
		return models.Position{}, err
	}
	return position, nil
}

func (r *position) FindByID(ctx context.Context, id int) (models.Position, error) {
	var position models.Position
	err := r.db.Select("id", "name", "base_salary").Where("id = ?", id).First(&position).Error
	if err != nil {
		return models.Position{}, err
	}
	return position, nil
}
func (r *position) Update(ctx context.Context, id int, position models.Position) (models.Position, error) {
	err := r.db.Model(&models.Position{}).Where("id = ?", id).Updates(position).Error
	if err != nil {
		return models.Position{}, err
	}
	return position, nil
}
func (r *position) Delete(ctx context.Context, id int) error {
	err := r.db.Where("id = ?", id).Delete(&models.Position{}).Error
	if err != nil {
		return err
	}
	return nil
}
func (r *position) FindAll(ctx context.Context, base dto.BaseQuery) (positions []models.Position, totalData int64, err error) {
	err = r.db.Scopes(utils.Paginate(base.Page, base.Limit),
		utils.Search(base.Search)).Select("name", "base_salary").Find(&positions).Error

	if err != nil {
		return nil, 0, err
	}

	err = r.db.Model(&models.Position{}).Scopes(utils.Search(base.Search)).Count(&totalData).Error

	if err != nil {
		return nil, 0, err
	}
	return positions, totalData, nil

}

func (r *position) BulkCreate(ctx context.Context, positions []models.Position) ([]models.Position, error) {
	err := r.db.Create(&positions).Error
	if err != nil {
		return nil, err
	}
	return positions, nil
}

func (r *position) FindAllOnly(ctx context.Context) (positions []models.Position, err error) {
	err = r.db.Select("id", "name").Find(&positions).Error
	if err != nil {
		return nil, err
	}
	return positions, nil
}
