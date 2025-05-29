package repositories

import (
	"context"

	"github.com/handarudwiki/payroll-sistem/internal/dto"
	"github.com/handarudwiki/payroll-sistem/internal/models"
	"github.com/handarudwiki/payroll-sistem/internal/utils"
	"gorm.io/gorm"
)

type (
	Attendance interface {
		FindByID(ctx context.Context, id int) (models.Attendance, error)
		FindAll(ctx context.Context, base dto.BaseQuery) (attendances []models.Attendance, totalData int64, err error)
		Create(ctx context.Context, attendance models.Attendance) (models.Attendance, error)
		Update(ctx context.Context, id int, attendance models.Attendance) (models.Attendance, error)
		Delete(ctx context.Context, id int) error
		TodayAttendance(ctx context.Context) ([]models.Attendance, error)
		BulkCreate(ctx context.Context, attendances []models.Attendance) ([]models.Attendance, error)
	}

	attendance struct {
		db *gorm.DB
	}
)

func NewAttendanceRepository(db *gorm.DB) Attendance {
	return &attendance{
		db: db,
	}
}

func (r *attendance) FindByID(ctx context.Context, id int) (models.Attendance, error) {
	var attendance models.Attendance
	err := r.db.Where("id = ?", id).First(&attendance).Error
	if err != nil {
		return models.Attendance{}, err
	}
	return attendance, nil
}

func (r *attendance) FindAll(ctx context.Context, base dto.BaseQuery) (attendances []models.Attendance, totalData int64, err error) {
	err = r.db.Model(&models.Attendance{}).Scopes(
		utils.Paginate(base.Page, base.Limit),
		utils.Search(base.Search),
	).Find(&attendances).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.Model(&models.Attendance{}).Scopes(utils.Search(base.Search)).Count(&totalData).Error
	if err != nil {
		return nil, 0, err
	}

	return attendances, totalData, nil
}

func (r *attendance) Create(ctx context.Context, attendance models.Attendance) (models.Attendance, error) {
	err := r.db.Create(&attendance).Error
	if err != nil {
		return models.Attendance{}, err
	}
	return attendance, nil
}

func (r *attendance) Update(ctx context.Context, id int, attendance models.Attendance) (models.Attendance, error) {
	err := r.db.Model(&models.Attendance{}).Where("id = ?", id).Updates(attendance).Error
	if err != nil {
		return models.Attendance{}, err
	}
	return attendance, nil
}
func (r *attendance) Delete(ctx context.Context, id int) error {
	err := r.db.Where("id = ?", id).Delete(&models.Attendance{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *attendance) TodayAttendance(ctx context.Context) ([]models.Attendance, error) {
	var attendances []models.Attendance
	err := r.db.Where("date = NOW()").Find(&attendances).Error
	if err != nil {
		return nil, err
	}
	return attendances, nil
}

func (r *attendance) BulkCreate(ctx context.Context, attendances []models.Attendance) ([]models.Attendance, error) {
	err := r.db.Create(&attendances).Error
	if err != nil {
		return nil, err
	}
	return attendances, nil
}
