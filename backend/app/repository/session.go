package repository

import (
	"github.com/google/uuid"
	"sudoku/model"
)

type ISessionRepository interface {
	FindByID(id uuid.UUID) (*model.Session, error)
	Save(session *model.Session) error
}
