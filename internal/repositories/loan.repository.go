package repositories

import (
	"context"

	"github.com/handarudwiki/payroll-sistem/internal/dto"
	"github.com/handarudwiki/payroll-sistem/internal/models"
	"github.com/handarudwiki/payroll-sistem/internal/utils"
	"gorm.io/gorm"
)

type (
	Loan interface {
		FindByID(ctx context.Context, id int) (models.Loan, error)
		FindAll(ctx context.Context, base dto.BaseQuery) (loans []models.Loan, totalData int64, err error)
		Create(ctx context.Context, loan models.Loan) (models.Loan, error)
		Update(ctx context.Context, id int, loan models.Loan) (models.Loan, error)
		Delete(ctx context.Context, id int) error
		BulkCreate(ctx context.Context, loans []models.Loan) ([]models.Loan, error)
	}

	loan struct {
		db *gorm.DB
	}
)

func NewLoanRepository(db *gorm.DB) Loan {
	return &loan{
		db: db,
	}
}

func (r *loan) FindByID(ctx context.Context, id int) (models.Loan, error) {
	var loan models.Loan
	err := r.db.Where("id = ?", id).Preload("Employee").First(&loan).Error
	if err != nil {
		return models.Loan{}, err
	}
	return loan, nil
}
func (r *loan) FindAll(ctx context.Context, base dto.BaseQuery) (loans []models.Loan, totalData int64, err error) {
	err = r.db.Model(&models.Loan{}).Scopes(
		utils.Paginate(base.Page, base.Limit),
		utils.Search(base.Search),
	).Find(&loans).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.Model(&models.Loan{}).Scopes(utils.Search(base.Search)).Count(&totalData).Error
	if err != nil {
		return nil, 0, err
	}

	return loans, totalData, nil
}
func (r *loan) Create(ctx context.Context, loan models.Loan) (models.Loan, error) {
	err := r.db.Create(&loan).Error
	if err != nil {
		return models.Loan{}, err
	}
	return loan, nil
}
func (r *loan) Update(ctx context.Context, id int, loan models.Loan) (models.Loan, error) {
	err := r.db.Model(&models.Loan{}).Where("id = ?", id).Updates(loan).Error
	if err != nil {
		return models.Loan{}, err
	}
	return loan, nil
}
func (r *loan) Delete(ctx context.Context, id int) error {
	err := r.db.Where("id = ?", id).Delete(&models.Loan{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *loan) BulkCreate(ctx context.Context, loans []models.Loan) ([]models.Loan, error) {
	err := r.db.Create(&loans).Error
	if err != nil {
		return nil, err
	}
	return loans, nil
}
