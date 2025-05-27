package repositories

import (
	"context"

	"github.com/handarudwiki/payroll-sistem/internal/dto"
	"github.com/handarudwiki/payroll-sistem/internal/models"
	"github.com/handarudwiki/payroll-sistem/internal/utils"
	"gorm.io/gorm"
)

type (
	SalaryComponent interface {
		FindByID(ctx context.Context, id int) (models.SalaryComponent, error)
		FindAll(ctx context.Context, base dto.BaseQuery) (salaryComponents []models.SalaryComponent, totalData int64, err error)
		Create(ctx context.Context, salaryComponent models.SalaryComponent) (models.SalaryComponent, error)
		Update(ctx context.Context, id int, salaryComponent models.SalaryComponent) (models.SalaryComponent, error)
		Delete(ctx context.Context, id int) error
		BulkCreate(ctx context.Context, salaryComponents []models.SalaryComponent) ([]models.SalaryComponent, error)
		FindAllOnly(ctx context.Context) (salaryComponents []models.SalaryComponent, err error)
	}
	salaryComponent struct {
		db *gorm.DB
	}
)

func NewSalaryComponentRepository(db *gorm.DB) SalaryComponent {
	return &salaryComponent{
		db: db,
	}
}

func (r *salaryComponent) FindByID(ctx context.Context, id int) (models.SalaryComponent, error) {
	var salaryComponent models.SalaryComponent
	err := r.db.Where("id = ?", id).First(&salaryComponent).Error
	if err != nil {
		return models.SalaryComponent{}, err
	}
	return salaryComponent, nil
}

func (r *salaryComponent) FindAll(ctx context.Context, base dto.BaseQuery) (salaryComponents []models.SalaryComponent, totalData int64, err error) {
	err = r.db.Model(&models.SalaryComponent{}).Scopes(
		utils.Paginate(base.Page, base.Limit),
		utils.Search(base.Search),
	).Find(&salaryComponents).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.Model(&models.SalaryComponent{}).Scopes(utils.Search(base.Search)).Count(&totalData).Error
	if err != nil {
		return nil, 0, err
	}

	return salaryComponents, totalData, nil
}

func (r *salaryComponent) Create(ctx context.Context, salaryComponent models.SalaryComponent) (models.SalaryComponent, error) {
	err := r.db.Create(&salaryComponent).Error
	if err != nil {
		return models.SalaryComponent{}, err
	}
	return salaryComponent, nil
}
func (r *salaryComponent) Update(ctx context.Context, id int, salaryComponent models.SalaryComponent) (models.SalaryComponent, error) {
	err := r.db.Model(&models.SalaryComponent{}).Where("id = ?", id).Updates(salaryComponent).Update("is_recurring", salaryComponent.IsRecurring).Error
	if err != nil {
		return models.SalaryComponent{}, err
	}
	return salaryComponent, nil
}
func (r *salaryComponent) Delete(ctx context.Context, id int) error {
	err := r.db.Where("id = ?", id).Delete(&models.SalaryComponent{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *salaryComponent) BulkCreate(ctx context.Context, salaryComponents []models.SalaryComponent) ([]models.SalaryComponent, error) {
	err := r.db.Create(&salaryComponents).Error
	if err != nil {
		return nil, err
	}
	return salaryComponents, nil
}
func (r *salaryComponent) FindAllOnly(ctx context.Context) (salaryComponents []models.SalaryComponent, err error) {
	err = r.db.Find(&salaryComponents).Error
	if err != nil {
		return nil, err
	}
	return salaryComponents, nil
}
