package models

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"time"
)

type BaseModel struct {
	ID uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`

	// Date auto generation by gorm
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// Auto generate uuid on new record
func (BaseModel) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("ID", uuid.NewV4())
}
