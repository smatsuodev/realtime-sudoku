package gorm

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type UUIDBaseModel struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type User struct {
	gorm.Model
	Name         string `gorm:"not null"`
	GitHubUserID uint   `gorm:"column:github_user_id;unique;not null"`
	Sessions     []Session
}

type Session struct {
	UUIDBaseModel
	UserID    uint `gorm:"not null"`
	User      User
	ExpiresAt time.Time `gorm:"not null"`
}
