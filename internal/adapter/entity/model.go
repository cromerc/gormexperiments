package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type ModelUUID struct {
	ID        uuid.UUID `gorm:"<-:create;type:uuid;primary_key;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (m *ModelUUID) BeforeCreate(*gorm.DB) (err error) {
	m.ID = uuid.New()
	return
}

type ModelID struct {
	ID        int64 `gorm:"<-:create;primary_key;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
