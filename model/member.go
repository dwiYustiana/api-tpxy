package model

import (
	"api-tpx/model/entity"

	"github.com/jinzhu/gorm"
)

type (
	MemberModelInterface interface {
		GetMemberById(id string) (*entity.MemberEntity, error)
		GetAllMember() ([]*entity.MemberEntity, error)
	}
	memberModel struct {
		DB *gorm.DB
	}
)

func init() {

}

func NewMemberModel(db *gorm.DB) MemberModelInterface {
	return &memberModel{db}
}
func (m *memberModel) GetMemberById(id string) (*entity.MemberEntity, error) {
	Member := entity.MemberEntity{}
	if err := m.DB.Where("id = ?", id).First(&Member).Error; err != nil {
		return &Member, err
	}
	return &Member, nil
}
func (m *memberModel) GetAllMember() ([]*entity.MemberEntity, error) {
	Member := make([]*entity.MemberEntity, 0)
	if err := m.DB.Table("Members").Find(&Member).Error; err != nil {
		return Member, err
	}
	return Member, nil
}
