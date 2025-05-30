package repositories

import (
	"context"

	"github.com/handarudwiki/payroll-sistem/internal/dto"
	"github.com/handarudwiki/payroll-sistem/internal/models"
	"github.com/handarudwiki/payroll-sistem/internal/utils"
	"gorm.io/gorm"
)

type (
	Employee interface {
		FindByID(ctx context.Context, id int) (models.Employee, error)
		FindByEmail(ctx context.Context, email string) (models.Employee, error)
		FindByNIK(ctx context.Context, nik string) (models.Employee, error)
		FindByPhone(ctx context.Context, phone string) (models.Employee, error)
		Create(ctx context.Context, employee models.Employee) (models.Employee, error)
		Update(ctx context.Context, id int, employee models.Employee) (models.Employee, error)
		Delete(ctx context.Context, id int) error
		FindAll(ctx context.Context, base dto.BaseQuery) (employees []models.Employee, totalData int64, err error)
		FindAllActive(ctx context.Context) (employees []models.Employee, err error)
		FindByIDSActive(ctx context.Context, ids []int) ([]models.Employee, error)
		BulkCreate(ctx context.Context, employees []models.Employee) ([]models.Employee, error)
	}

	employee struct {
		db *gorm.DB
	}
)

func NewEmployeeRepository(db *gorm.DB) Employee {
	return &employee{
		db: db,
	}
}

func (r *employee) FindByID(ctx context.Context, id int) (models.Employee, error) {
	var employee models.Employee
	err := r.db.Where("id = ?", id).Preload("Department").
		Preload("Position").First(&employee).First(&employee).Error
	if err != nil {
		return models.Employee{}, err
	}
	return employee, nil
}

func (r *employee) FindByEmail(ctx context.Context, email string) (models.Employee, error) {
	var employee models.Employee
	err := r.db.Where("email = ?", email).First(&employee).Error
	if err != nil {
		return models.Employee{}, err
	}
	return employee, nil
}

func (r *employee) FindByNIK(ctx context.Context, nik string) (models.Employee, error) {
	var employee models.Employee
	err := r.db.Where("nik = ?", nik).First(&employee).Error
	if err != nil {
		return models.Employee{}, err
	}
	return employee, nil
}

func (r *employee) FindByPhone(ctx context.Context, phone string) (models.Employee, error) {
	var employee models.Employee
	err := r.db.Where("phone = ?", phone).First(&employee).Error
	if err != nil {
		return models.Employee{}, err
	}
	return employee, nil
}

func (r *employee) Create(ctx context.Context, employee models.Employee) (models.Employee, error) {
	err := r.db.Create(&employee).Error
	if err != nil {
		return models.Employee{}, err
	}
	return employee, nil
}

func (r *employee) Update(ctx context.Context, id int, employee models.Employee) (models.Employee, error) {
	err := r.db.Model(&models.Employee{}).Where("id = ?", id).Updates(employee).Error
	if err != nil {
		return models.Employee{}, err
	}
	return employee, nil
}
func (r *employee) Delete(ctx context.Context, id int) error {
	err := r.db.Where("id = ?", id).Delete(&models.Employee{}).Error
	if err != nil {
		return err
	}
	return nil
}
func (r *employee) FindAll(ctx context.Context, base dto.BaseQuery) (employees []models.Employee, totalData int64, err error) {
	err = r.db.Scopes(utils.Paginate(base.Page, base.Limit),
		utils.Search(base.Search), utils.FilterDepartment(base.DepartmentID), utils.FilterPosition(base.PositionID)).Preload("Department").
		Preload("Position").Find(&employees).Error
	if err != nil {
		return employees, 0, err
	}
	err = r.db.Model(&models.Employee{}).Scopes(utils.Search(base.Search)).Count(&totalData).Error
	if err != nil {
		return employees, 0, err
	}
	return employees, totalData, nil
}

func (r *employee) FindAllActive(ctx context.Context) (employees []models.Employee, err error) {
	err = r.db.Select("id", "name", "nik").
		Where("status", "active").
		Preload("Department").Preload("Position").
		Preload("Leaves").
		Preload("EmployeeComponent.SalaryComponent").
		Preload("Loans").Preload("Attendances").
		Find(&employees).Error
	if err != nil {
		return nil, err
	}
	return employees, nil
}

func (r *employee) FindByIDSActive(ctx context.Context, ids []int) ([]models.Employee, error) {
	var employees []models.Employee
	err := r.db.Select("id", "name", "nik").
		Where("status = ? AND id IN ?", "active", ids).
		Preload("Department").Preload("Position").
		Preload("EmployeeComponent.SalaryComponent").
		Preload("Leaves").Preload("Loans").Preload("Attendances").
		Find(&employees).Error
	if err != nil {
		return nil, err
	}
	return employees, nil
}

func (r *employee) BulkCreate(ctx context.Context, employees []models.Employee) ([]models.Employee, error) {
	err := r.db.Create(&employees).Error
	if err != nil {
		return nil, err
	}
	return employees, nil
}
