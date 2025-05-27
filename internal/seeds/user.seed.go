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
	}

	_, err = repo.BulkCreate(context.Background(), users)
	if err != nil {
		return err
	}

	return

}
