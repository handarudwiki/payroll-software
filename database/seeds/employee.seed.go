package seeds

import (
	"context"

	"github.com/handarudwiki/payroll-sistem/internal/models"
	"github.com/handarudwiki/payroll-sistem/internal/repositories"
)

func EmployeeSeed(employeeRepo repositories.Employee, departmentRepo repositories.Department, positionRepo repositories.Position, userRepo repositories.User) (err error) {
	departments, err := departmentRepo.FindAllOnly(context.Background())
	if err != nil {
		return err
	}
	positions, err := positionRepo.FindAllOnly(context.Background())
	if err != nil {
		return err
	}
	users, err := userRepo.FindAllOnly(context.Background())
	if err != nil {
		return err
	}

	employees := []models.Employee{
		{
			Name:         "Handarud Wiki",
			Email:        "handaru@gmail.com",
			Phone:        "08123456789",
			NIK:          "1234567890",
			DepartmentID: departments[0].ID,
			PositionID:   positions[0].ID,
			UserID:       users[0].ID,
			HireDate:     "2023-01-01",
			Status:       models.STATUS_ACTIVE,
		},
		{
			Name:         "Siti Maesaroh",
			Email:        "siti.maesaroh@gmail.com",
			Phone:        "08123456780",
			NIK:          "1234567891",
			DepartmentID: departments[1].ID,
			PositionID:   positions[1].ID,
			UserID:       users[1].ID,
			HireDate:     "2023-02-01",
			Status:       models.STATUS_ACTIVE,
		},
		{
			Name:         "Ahmad Yani",
			Email:        "ahmad.yani@gmail.com",
			Phone:        "08123456781",
			NIK:          "1234567892",
			DepartmentID: departments[2].ID,
			PositionID:   positions[2].ID,
			UserID:       users[2].ID,
			HireDate:     "2023-03-01",
			Status:       models.STATUS_ACTIVE,
		},
		{
			Name:         "Nur Aisyah",
			Email:        "nur.aisyah@gmail.com",
			Phone:        "08123456782",
			NIK:          "1234567893",
			DepartmentID: departments[3].ID,
			PositionID:   positions[3].ID,
			UserID:       users[3].ID,
			HireDate:     "2023-04-01",
			Status:       models.STATUS_ACTIVE,
		},
		{
			Name:         "Dedi Supriyadi",
			Email:        "dedi.supriyadi@gmail.com",
			Phone:        "08123456783",
			NIK:          "1234567894",
			DepartmentID: departments[4].ID,
			PositionID:   positions[4].ID,
			UserID:       users[4].ID,
			HireDate:     "2023-05-01",
			Status:       models.STATUS_ACTIVE,
		},
		{
			Name:         "Indah Permatasari",
			Email:        "indah.permatasari@gmail.com",
			Phone:        "08123456784",
			NIK:          "1234567895",
			DepartmentID: departments[5].ID,
			PositionID:   positions[5].ID,
			UserID:       users[5].ID,
			HireDate:     "2023-06-01",
			Status:       models.STATUS_ACTIVE,
		},
		{
			Name:         "Yusuf Maulana",
			Email:        "yusuf.maulana@gmail.com",
			Phone:        "08123456785",
			NIK:          "1234567896",
			DepartmentID: departments[6].ID,
			PositionID:   positions[6].ID,
			UserID:       users[6].ID,
			HireDate:     "2023-07-01",
			Status:       models.STATUS_ACTIVE,
		},
		{
			Name:         "Rina Oktaviani",
			Email:        "rina.oktaviani@gmail.com",
			Phone:        "08123456786",
			NIK:          "1234567897",
			DepartmentID: departments[7].ID,
			PositionID:   positions[7].ID,
			UserID:       users[7].ID,
			HireDate:     "2023-08-01",
			Status:       models.STATUS_ACTIVE,
		},
		{
			Name:         "Galih Prasetya",
			Email:        "galih.prasetya@gmail.com",
			Phone:        "08123456787",
			NIK:          "1234567898",
			DepartmentID: departments[8].ID,
			PositionID:   positions[8].ID,
			UserID:       users[8].ID,
			HireDate:     "2023-09-01",
			Status:       models.STATUS_ACTIVE,
		},
		{
			Name:         "Yuliana Dewi",
			Email:        "yuliana.dewi@gmail.com",
			Phone:        "08123456788",
			NIK:          "1234567899",
			DepartmentID: departments[9].ID,
			PositionID:   positions[9].ID,
			UserID:       users[9].ID,
			HireDate:     "2023-10-01",
			Status:       models.STATUS_ACTIVE,
		},
		{
			Name:         "Rizky Ramadhan",
			Email:        "rizky.ramadhan@gmail.com",
			Phone:        "08123456790",
			NIK:          "1234567800",
			DepartmentID: departments[10].ID,
			PositionID:   positions[10].ID,
			UserID:       users[10].ID,
			HireDate:     "2023-11-01",
			Status:       models.STATUS_ACTIVE,
		},
	}

	_, err = employeeRepo.BulkCreate(context.Background(), employees)
	if err != nil {
		return err
	}
	return nil

}
