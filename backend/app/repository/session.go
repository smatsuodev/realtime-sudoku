package repository

import (
	"github.com/google/uuid"
	"sudoku/model"
)

type ISessionRepository interface {
	FindByID(id uuid.UUID) (*model.Session, error)
	// Save 新規レコード作成の場合, 採番された ID をセットする
	Save(session *model.Session) error
}
