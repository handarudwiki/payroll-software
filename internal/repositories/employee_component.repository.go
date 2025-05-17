package repositories

import (
	"context"

	"github.com/handarudwiki/payroll-sistem/internal/dto"
	"github.com/handarudwiki/payroll-sistem/internal/models"
	"github.com/handarudwiki/payroll-sistem/internal/utils"
	"gorm.io/gorm"
)

type (
	EmployeeComponent interface {
		Create(ctx context.Context, employeeComponent models.EmployeeComponent) (models.EmployeeComponent, error)
		FindByID(ctx context.Context, id int) (models.EmployeeComponent, error)
		Update(ctx context.Context, id int, employeeComponent models.EmployeeComponent) (models.EmployeeComponent, error)
		Delete(ctx context.Context, id int) error
		FindAll(ctx context.Context, base dto.BaseQuery) (employeeComponents []models.EmployeeComponent, totalData int64, err error)
	}

	employeeComponent struct {
		db *gorm.DB
	}
)

func NewEmployeeComponentRepository(db *gorm.DB) EmployeeComponent {
	return &employeeComponent{
		db: db,
	}
}

func (r *employeeComponent) Create(ctx context.Context, employeeComponent models.EmployeeComponent) (models.EmployeeComponent, error) {
	err := r.db.Create(&employeeComponent).Error
	if err != nil {
		return models.EmployeeComponent{}, err
	}
	return employeeComponent, nil
}

func (r *employeeComponent) FindByID(ctx context.Context, id int) (models.EmployeeComponent, error) {
	var employeeComponent models.EmployeeComponent
	err := r.db.Where("id = ?", id).First(&employeeComponent).Error
	if err != nil {
		return models.EmployeeComponent{}, err
	}
	return employeeComponent, nil
}

func (r *employeeComponent) Update(ctx context.Context, id int, employeeComponent models.EmployeeComponent) (models.EmployeeComponent, error) {
	err := r.db.Model(&models.EmployeeComponent{}).Where("id = ?", id).Updates(employeeComponent).Error
	if err != nil {
		return models.EmployeeComponent{}, err
	}
	return employeeComponent, nil
}

func (r *employeeComponent) Delete(ctx context.Context, id int) error {
	err := r.db.Where("id = ?", id).Delete(&models.EmployeeComponent{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *employeeComponent) FindAll(ctx context.Context, base dto.BaseQuery) (employeeComponents []models.EmployeeComponent, totalData int64, err error) {
	err = r.db.Model(&models.EmployeeComponent{}).Scopes(
		utils.Paginate(base.Page, base.Limit),
		utils.Search(base.Search),
	).Find(&employeeComponents).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.Model(&models.EmployeeComponent{}).Scopes(utils.Search(base.Search)).Count(&totalData).Error
	if err != nil {
		return nil, 0, err
	}

	return employeeComponents, totalData, nil
}
