//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE
package repository

import (
	"github.com/google/uuid"
	"github.com/samber/mo"
	"sudoku/model"
)

type ISessionRepository interface {
	FindByID(id uuid.UUID) (mo.Option[*model.Session], error)
	// Save 新規レコード作成の場合, 採番された ID をセットする
	Save(session *model.Session) error
	DeleteExpired() error
}
