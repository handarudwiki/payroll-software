package repositories

import (
	"context"

	"github.com/handarudwiki/payroll-sistem/internal/dto"
	"github.com/handarudwiki/payroll-sistem/internal/models"
	"github.com/handarudwiki/payroll-sistem/internal/utils"
	"gorm.io/gorm"
)

type (
	Department interface {
		Create(ctx context.Context, department models.Department) (models.Department, error)
		FindByID(ctx context.Context, id int) (models.Department, error)
		Update(ctx context.Context, id int, department models.Department) (models.Department, error)
		Delete(ctx context.Context, id int) error
		FindAll(ctx context.Context, base dto.BaseQuery) (departments []models.Department, totalData int64, err error)
		BulkCreate(ctx context.Context, departments []models.Department) ([]models.Department, error)
		FindAllOnly(ctx context.Context) (departments []models.Department, err error)
	}
	DepartmentRepository struct {
		db *gorm.DB
	}
)

func NewDepartmentRepository(db *gorm.DB) Department {
	return &DepartmentRepository{
		db: db,
	}

}

func (r *DepartmentRepository) Create(ctx context.Context, department models.Department) (models.Department, error) {
	err := r.db.Create(&department).Error
	if err != nil {
		return models.Department{}, err
	}
	return department, nil
}

func (r *DepartmentRepository) FindByID(ctx context.Context, id int) (models.Department, error) {
	var department models.Department
	err := r.db.Select("id", "name", "description").Where("id = ?", id).First(&department).Error
	if err != nil {
		return models.Department{}, err
	}
	return department, nil
}

func (r *DepartmentRepository) Update(ctx context.Context, id int, department models.Department) (models.Department, error) {
	err := r.db.Model(&models.Department{}).Where("id = ?", id).Updates(department).Error
	if err != nil {
		return models.Department{}, err
	}
	return department, nil
}

func (r *DepartmentRepository) Delete(ctx context.Context, id int) error {
	err := r.db.Where("id = ?", id).Delete(&models.Department{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *DepartmentRepository) FindAll(ctx context.Context, base dto.BaseQuery) (departments []models.Department, totalData int64, err error) {
	err = r.db.Scopes(utils.Paginate(base.Page, base.Limit),
		utils.Search(base.Search)).Select("id", "name", "description").
		Find(&departments).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.Model(&models.Department{}).Scopes(utils.Search(base.Search)).Count(&totalData).Error

	if err != nil {
		return nil, 0, err
	}
	return departments, totalData, nil
}

func (r *DepartmentRepository) BulkCreate(ctx context.Context, departments []models.Department) ([]models.Department, error) {
	err := r.db.Create(&departments).Error
	if err != nil {
		return nil, err
	}
	return departments, nil
}

func (r *DepartmentRepository) FindAllOnly(ctx context.Context) (departments []models.Department, err error) {
	err = r.db.Select("id", "name").Find(&departments).Error
	if err != nil {
		return nil, err
	}
	return departments, nil
}
