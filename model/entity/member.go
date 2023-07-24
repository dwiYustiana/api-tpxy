package entity

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type MemberEntity struct {
	ID        string    `gorm:"type:varchar(50);primary_key;column:id"`
	Nik       string    `gorm:"type:varchar(50);column:nik";unique_index`
	FullName  string    `gorm:"type:varchar(50);column:full_name"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

// TableName ...
func (MemberEntity) TableName() string {
	return "members"
}

func (item *MemberEntity) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	return scope.SetColumn("ID", uuid.String())
}
