package request

type PurchaseCreditRequest struct {
	LimitCode        string `json:"limit_code" validate:"required"`
	TotalDebitAmount int    `json:"debit_amount" validate:"required"`
}
