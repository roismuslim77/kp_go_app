package transaction

type CheckoutLoanReq struct {
	OTRPrice      float64 `json:"otr_price" binding:"required"`
	AdminFee      float64 `json:"admin_fee" binding:"required"`
	Tenor         int     `json:"tenor" binding:"required"`
	InterestPrice float64 `json:"interest_price" binding:"required"`
	Name          string  `json:"name" binding:"required"`
}
