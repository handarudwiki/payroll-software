package models

import (
	"time"

	"github.com/handarudwiki/payroll-sistem/internal/dto"
)

type Position struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	BaseSalary float64   `json:"base_salary"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func NewPositionFromCreatePosition(data dto.CreatePosition) Position {
	return Position{
		Name:       data.Name,
		BaseSalary: data.BaseSalary,
	}
}

func NewPositionFromUpdatePosition(data dto.UpdatePosition) Position {
	return Position{
		Name:       data.Name,
		BaseSalary: data.BaseSalary,
	}
}
