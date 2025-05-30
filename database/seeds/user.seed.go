package seeds

import (
	"context"

	"github.com/handarudwiki/payroll-sistem/internal/models"
	"github.com/handarudwiki/payroll-sistem/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

func UserSeed(repo repositories.User) (err error) {

	password, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	users := []models.User{
		{
			Name:     "Admin",
			Username: "adminsudrajad",
			Password: string(password),
			Role:     models.RoleAdmin,
		},
		{
			Name:     "John Doe",
			Username: "johndoe",
			Password: string(password),
			Role:     models.RoleUser,
		},
		{
			Name:     "Jane Smith",
			Username: "janesmith",
			Password: string(password),
			Role:     models.RoleUser,
		},
		{
			Name:     "Michael Johnson",
			Username: "mjohnson",
			Password: string(password),
			Role:     models.RoleUser,
		},
		{
			Name:     "Emily Davis",
			Username: "edavis",
			Password: string(password),
			Role:     models.RoleUser,
		},
		{
			Name:     "David Lee",
			Username: "dlee",
			Password: string(password),
			Role:     models.RoleUser,
		},
		{
			Name:     "Sarah Kim",
			Username: "skim",
			Password: string(password),
			Role:     models.RoleUser,
		},
		{
			Name:     "James Brown",
			Username: "jbrown",
			Password: string(password),
			Role:     models.RoleUser,
		},
		{
			Name:     "Linda White",
			Username: "lwhite",
			Password: string(password),
			Role:     models.RoleUser,
		},
		{
			Name:     "Robert Harris",
			Username: "rharris",
			Password: string(password),
			Role:     models.RoleUser,
		},
		{
			Name:     "Karen Martinez",
			Username: "kmartinez",
			Password: string(password),
			Role:     models.RoleUser,
		},
	}

	_, err = repo.BulkCreate(context.Background(), users)
	if err != nil {
		return err
	}

	return

}
