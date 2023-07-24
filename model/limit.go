package model

import (
	"api-tpx/model/entity"

	"github.com/jinzhu/gorm"
)

type (
	LimitModelInterface interface {
		GetLimitById(id string) *entity.LimitEntity
		GetLimit(memberCode string) ([]*entity.LimitEntity, error)
	}
	limitModel struct {
		DB *gorm.DB
	}
)

func init() {

}

func NewLimitModel(db *gorm.DB) LimitModelInterface {
	return &limitModel{db}
}
func (m *limitModel) GetLimitById(id string) *entity.LimitEntity {
	limit := entity.LimitEntity{}
	if err := m.DB.Where("id = ?", id).First(&limit).Error; err != nil {
	}
	return &limit
}
func (m *limitModel) GetLimit(memberCode string) ([]*entity.LimitEntity, error) {
	limit := make([]*entity.LimitEntity, 0)
	if err := m.DB.Where("member_code = ?", memberCode).Find(&limit).Error; err != nil {
		return limit, err
	}
	return limit, nil
}
