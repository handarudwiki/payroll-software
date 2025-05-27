package repositories

import (
	"context"

	"github.com/handarudwiki/payroll-sistem/internal/models"
	"gorm.io/gorm"
)

type (
	Payroll interface {
		Create(ctx context.Context, tx *gorm.DB, payroll models.Payroll) (models.Payroll, error)
		FindByID(ctx context.Context, id int) (models.Payroll, error)
	}

	payroll struct {
		db *gorm.DB
	}
)

func NewPayrollRepository(db *gorm.DB) Payroll {
	return &payroll{
		db: db,
	}
}

func (r *payroll) Create(ctx context.Context, tx *gorm.DB, payroll models.Payroll) (models.Payroll, error) {
	if tx == nil {
		tx = r.db.WithContext(ctx)
	}

	err := tx.Create(&payroll).Error
	if err != nil {
		return models.Payroll{}, err
	}
	return payroll, nil
}

func (r *payroll) FindByID(ctx context.Context, id int) (models.Payroll, error) {
	var payroll models.Payroll
	err := r.db.Where("id = ?", id).Preload("Employee").Preload("PayslipDetail.SalaryComponent").First(&payroll).Error
	if err != nil {
		return models.Payroll{}, err
	}
	return payroll, nil
}
