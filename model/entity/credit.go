package entity

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type CreditEntity struct {
	ID             string    `gorm:"type:varchar(50);primary_key;column:id"`
	MemberCode     string    `gorm:"type:varchar(50);column:member_code"`
	LimitCode      string    `gorm:"type:varchar(11);column:limit_code"`
	RemainingLimit int       `gorm:"type:integer(11);column:amount"`
	CreatedAt      time.Time `gorm:"column:created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at"`
}

// TableName ...
func (CreditEntity) TableName() string {
	return "Credits"
}

func (item *CreditEntity) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	return scope.SetColumn("ID", uuid.String())
}
