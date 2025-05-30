package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/handarudwiki/payroll-sistem/config"
	"github.com/handarudwiki/payroll-sistem/internal/dto"
	"github.com/handarudwiki/payroll-sistem/internal/models"
	"github.com/handarudwiki/payroll-sistem/internal/models/commons"
	"github.com/handarudwiki/payroll-sistem/internal/repositories"
	"github.com/handarudwiki/payroll-sistem/internal/responses"
	"github.com/handarudwiki/payroll-sistem/internal/utils"
	"gorm.io/gorm"
)

type (
	User interface {
		Login(ctx context.Context, dto dto.Login) (responses.LoginResponse, error)
		Register(ctx context.Context, dto dto.Register) (models.UserResponse, error)
		Update(ctx context.Context, id int, dto dto.UpdateUser) (models.UserResponse, error)
		UpdatePassword(ctx context.Context, id int, dto dto.UpdatePassword) (models.UserResponse, error)
		Me(ctx context.Context, id int) (models.UserResponse, error)
	}

	userService struct {
		userRepository repositories.User
		jwt            config.JWT
	}
)

func NewUserService(userRepository repositories.User, jwt config.JWT) User {
	return &userService{
		userRepository: userRepository,
		jwt:            jwt,
	}
}
func (s *userService) Login(ctx context.Context, dto dto.Login) (res responses.LoginResponse, err error) {
	user, err := s.userRepository.FindByUsername(ctx, dto.Username)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return responses.LoginResponse{}, commons.ErrCredentials
	}

	if err != nil {
		return responses.LoginResponse{}, err
	}

	isPasswordMatch, err := user.ComparePassword(dto.Password)
	if err != nil {
		return responses.LoginResponse{}, err
	}

	if !isPasswordMatch {
		return responses.LoginResponse{}, commons.ErrCredentials
	}

	fmt.Println("role generate token ", user.Role)

	token, err := utils.GenerateToken(user.ID, user.Role, s.jwt.Secret)
	if err != nil {
		return responses.LoginResponse{}, err
	}

	return responses.LoginResponse{
		Token: token,
	}, nil
}

func (s *userService) Register(ctx context.Context, dto dto.Register) (res models.UserResponse, err error) {
	user, err := s.userRepository.FindByUsername(ctx, dto.Username)
	if err == nil {
		return models.UserResponse{}, err
	}

	if user.ID != 0 {
		return models.UserResponse{}, commons.ErrConflict
	}

	newUser := models.User{
		Name:     dto.Name,
		Username: dto.Username,
		Password: dto.Password,
		Role:     models.RoleUser,
	}
	err = newUser.EncryptPassword()

	if err != nil {
		return models.UserResponse{}, err
	}

	user, err = s.userRepository.Create(ctx, newUser)
	if err != nil {
		return models.UserResponse{}, err
	}

	return models.NewToUserResponse(user), nil
}

func (s *userService) Update(ctx context.Context, id int, dto dto.UpdateUser) (res models.UserResponse, err error) {
	user, err := s.userRepository.FindByID(ctx, id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.UserResponse{}, commons.ErrNotfound
	}

	if err != nil {
		return models.UserResponse{}, err
	}

	user.Name = dto.Name
	user.Username = dto.Username

	user, err = s.userRepository.Update(ctx, user, id)
	if err != nil {
		return models.UserResponse{}, err
	}

	return models.NewToUserResponse(user), nil
}
func (s *userService) UpdatePassword(ctx context.Context, id int, dto dto.UpdatePassword) (res models.UserResponse, err error) {
	user, err := s.userRepository.FindByID(ctx, id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.UserResponse{}, commons.ErrNotfound
	}

	if err != nil {
		return models.UserResponse{}, err
	}

	isPasswordMatch, err := user.ComparePassword(dto.OldPassword)
	if !isPasswordMatch {
		return models.UserResponse{}, commons.ErrWrongPassword
	}
	if err != nil {
		return models.UserResponse{}, err
	}

	err = user.EncryptPassword()
	if err != nil {
		return models.UserResponse{}, err
	}

	user, err = s.userRepository.UpdatePassword(ctx, id, user.Password)
	if err != nil {
		return models.UserResponse{}, err
	}

	return models.NewToUserResponse(user), nil
}

func (s *userService) Me(ctx context.Context, id int) (res models.UserResponse, err error) {
	user, err := s.userRepository.FindByID(ctx, id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.UserResponse{}, commons.ErrNotfound
	}

	if err != nil {
		return models.UserResponse{}, err
	}

	return models.NewToUserResponse(user), nil
}
