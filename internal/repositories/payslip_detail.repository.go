package repositories

import (
	"context"

	"github.com/handarudwiki/payroll-sistem/internal/models"
	"gorm.io/gorm"
)

type (
	PayslipDetail interface {
		Create(ctx context.Context, tx *gorm.DB, payslipDetail models.PayslipDetail) (models.PayslipDetail, error)
		BulkCreate(ctx context.Context, tx *gorm.DB, payslipDetails []models.PayslipDetail) error
	}

	payslipDetail struct {
		db *gorm.DB
	}
)

func NewPayslipDetailRepository(db *gorm.DB) PayslipDetail {
	return &payslipDetail{
		db: db,
	}
}

func (r *payslipDetail) Create(ctx context.Context, tx *gorm.DB, payslipDetail models.PayslipDetail) (models.PayslipDetail, error) {
	if tx == nil {
		tx = r.db.WithContext(ctx)
	}

	err := tx.Create(&payslipDetail).Error
	if err != nil {
		return models.PayslipDetail{}, err
	}
	return payslipDetail, nil
}

func (r *payslipDetail) BulkCreate(ctx context.Context, tx *gorm.DB, payslipDetails []models.PayslipDetail) error {
	if len(payslipDetails) == 0 {
		return nil // No data to insert
	}

	if tx == nil {
		tx = r.db.WithContext(ctx)
	}

	err := tx.Create(&payslipDetails).Error
	if err != nil {
		return err
	}
	return nil
}
