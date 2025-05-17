package responses

import "github.com/handarudwiki/payroll-sistem/internal/models"

type Position struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	BaseSalary float64 `json:"base_salary"`
}

func NewPositionResponse(position models.Position) Position {
	return Position{
		ID:         position.ID,
		Name:       position.Name,
		BaseSalary: position.BaseSalary,
	}
}

func NewPositionsResponse(positions []models.Position) []Position {
	var positionResponses []Position
	for _, position := range positions {
		positionResponses = append(positionResponses, NewPositionResponse(position))
	}
	return positionResponses
}
