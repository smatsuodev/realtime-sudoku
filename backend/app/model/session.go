package model

import (
	"github.com/google/uuid"
	"github.com/samber/mo"
	"time"
)

type Session struct {
	id        mo.Option[uuid.UUID]
	userID    uint
	expiresAt time.Time
}

func (s *Session) IsIDPresent() bool {
	return s.id.IsPresent()
}

func (s *Session) ID() uuid.UUID {
	if s.id.IsAbsent() {
		// Session 作成時のみ None になる
		return uuid.Nil
	}
	return s.id.MustGet()
}

func (s *Session) SetID(id uuid.UUID) {
	s.id = mo.Some(id)
}

func (s *Session) UserID() uint {
	return s.userID
}

func (s *Session) ExpiresAt() time.Time {
	return s.expiresAt
}

func NewSessionWithoutID(userID uint, expiresAt time.Time) *Session {
	return &Session{userID: userID, expiresAt: expiresAt}
}

func NewSession(id uuid.UUID, userID uint, expiresAt time.Time) *Session {
	return &Session{id: mo.Some(id), userID: userID, expiresAt: expiresAt}
}
