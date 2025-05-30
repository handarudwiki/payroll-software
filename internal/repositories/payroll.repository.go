package repositories

import (
	"context"

	"github.com/handarudwiki/payroll-sistem/internal/dto"
	"github.com/handarudwiki/payroll-sistem/internal/models"
	"github.com/handarudwiki/payroll-sistem/internal/utils"
	"gorm.io/gorm"
)

type (
	Payroll interface {
		Create(ctx context.Context, tx *gorm.DB, payroll models.Payroll) (models.Payroll, error)
		FindByID(ctx context.Context, id int) (models.Payroll, error)
		FindAll(ctx context.Context, base dto.BaseQuery) (payrolls []models.Payroll, totalData int64, err error)
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
	err := r.db.Where("id = ?", id).Preload("Employee.Department").Preload("Employee.Position").First(&payroll).Error
	if err != nil {
		return models.Payroll{}, err
	}
	return payroll, nil
}

func (r *payroll) FindAll(ctx context.Context, base dto.BaseQuery) (payrolls []models.Payroll, totalData int64, err error) {
	err = r.db.Model(&models.Payroll{}).Scopes(
		utils.Paginate(base.Page, base.Limit),
		utils.Search(base.Search),
		utils.FilterEmployeeID(base.EmployeeID),
		utils.FilterPeriodMonth(base.Period),
	).Preload("Employee.Department").Preload("Employee.Position").
		Find(&payrolls).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.Model(&models.Leave{}).Scopes(utils.Search(base.Search)).Count(&totalData).Error
	if err != nil {
		return nil, 0, err
	}

	return payrolls, totalData, nil
}
