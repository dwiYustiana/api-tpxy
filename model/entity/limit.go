package entity

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type LimitEntity struct {
	ID          string    `gorm:"type:varchar(50);primary_key;column:id"`
	MemberCode  string    `gorm:"type:varchar(50);column:member_code"`
	LimitAmount int       `gorm:"type:integer(11);column:limit_amount"`
	LimitTenor  int       `gorm:"type:integer(11);column:limit_tenor"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}

// TableName ...
func (LimitEntity) TableName() string {
	return "limits"
}

func (item *LimitEntity) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	return scope.SetColumn("ID", uuid.String())
}
