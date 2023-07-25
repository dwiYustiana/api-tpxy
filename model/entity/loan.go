package entity

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type LoanEntity struct {
	ID                string    `gorm:"type:varchar(50);primary_key;column:id"`
	LoanAmount        int       `gorm:"type:integer(11);column:loan_amount"`        //total pinjaman otr
	AdminFee          string    `gorm:"type:varchar(11);column:admin_fee"`          //admin fee
	AmountInstallment int       `gorm:"type:integer(11);column:amount_installment"` //jumlah cicilan
	InterestCredit    int       `gorm:"type:integer(11);column:interest_credit"`    //bunga cicilan
	ItemType          string    `gorm:"type:varchar(50);column:item_type"`          // type barang
	AssetName         string    `gorm:"type:varchar(50);column:asset_name"`         // asset Name
	Status            string    `gorm:"type:varchar(50);column:status"`             // status pinjaman
	CreatedAt         time.Time `gorm:"column:created_at"`
	UpdatedAt         time.Time `gorm:"column:updated_at"`
}

// TableName ...
func (LoanEntity) TableName() string {
	return "Loans"
}

func (item *LoanEntity) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	return scope.SetColumn("ID", uuid.String())
}
