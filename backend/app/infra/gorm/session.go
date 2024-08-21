package gorm

import (
	"errors"
	"github.com/google/uuid"
	"github.com/samber/mo"
	"gorm.io/gorm"
	"sudoku/model"
	"sudoku/repository"
)

var _ repository.ISessionRepository = (*SessionRepository)(nil)

type SessionRepository struct {
	db *gorm.DB
}

func NewSessionRepository(db *gorm.DB) *SessionRepository {
	return &SessionRepository{db: db}
}

func (r *SessionRepository) FindByID(id uuid.UUID) (mo.Option[*model.Session], error) {
	var session Session
	if err := r.db.First(&session, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return mo.None[*model.Session](), nil
		}
		return mo.None[*model.Session](), err
	}
	return mo.Some[*model.Session](
		model.NewSession(session.ID, session.UserID, session.ExpiresAt),
	), nil
}

func (r *SessionRepository) Save(session *model.Session) error {
	sessionModel := &Session{
		UserID:    session.UserID(),
		ExpiresAt: session.ExpiresAt(),
	}
	if session.IsIDPresent() {
		sessionModel.ID = session.ID()
	}

	if err := r.db.Save(sessionModel).Error; err != nil {
		return err
	}

	session.SetID(sessionModel.ID)

	return nil
}
