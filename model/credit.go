package model

import (
	"api-tpx/model/entity"

	"github.com/jinzhu/gorm"
)

type (
	CreditModelInterface interface {
		GetTotalCredit(memberCode string) *entity.CreditEntity
	}
	creditModel struct {
		DB *gorm.DB
	}
)

func init() {

}

func NewCreditModel(db *gorm.DB) LimitModelInterface {
	return &limitModel{db}
}
func (m *limitModel) GetRemainingLimit(memberCode, limitCode string) *entity.CreditEntity {

	limit := entity.CreditEntity{}
	if err := m.DB.Where("member_code = ? and limit_code = ?", memberCode).Find(&limit).Error; err != nil {
		return &limit
	}
	return &limit
}
