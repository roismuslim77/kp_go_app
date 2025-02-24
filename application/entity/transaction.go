package entity

import "time"

type Transaction struct {
	ID              int     `gorm:"column:id;type:int;primaryKey;autoIncrement:true;unique" json:"id"`
	CustomerId      int     `gorm:"column:customer_id;type:int" json:"customer_id"`
	CustomerLimitId int     `gorm:"column:customer_limit_id;type:int" json:"customer_limit_id"`
	ContractNumber  string  `gorm:"column:contract_number;type:string;size:255" json:"contract_number"`
	OTRPrice        float64 `gorm:"column:otr_price;type:float" json:"otr_price"`
	AdminFee        float64 `gorm:"column:admin_fee;type:float" json:"admin_fee"`
	InterestPrice   float64 `gorm:"column:interest_price;type:float" json:"interest_price"`
	Status          int     `gorm:"column:status;type:int" json:"status"`

	CreatedAt time.Time `gorm:"column:created_at;" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;" json:"updated_at"`

	Customer      Customer      `json:"customer" gorm:"foreignKey:CustomerId;references:ID"`
	CustomerLimit CustomerLimit `json:"customer_limit" gorm:"foreignKey:CustomerLimitId;references:ID"`

	// custom
	IsEmpty bool `gorm:"-" json:"-"`
}

func (t Transaction) TableName() string {
	return "transactions"
}
