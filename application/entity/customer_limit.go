package entity

import (
	"time"
)

type CustomerLimit struct {
	ID             int      `gorm:"column:id;type:int;primaryKey;autoIncrement:true;unique" json:"id"`
	CustomerId     int      `gorm:"column:customer_id;type:int" json:"customer_id"`
	Tenor          int      `gorm:"column:tenor;type:int" json:"tenor"`
	Limit          float64  `gorm:"column:limit;type:float" json:"limit"`
	RemainingLimit *float64 `gorm:"column:remaining_limit;type:float" json:"remaining_limit"`

	CreatedAt time.Time `gorm:"column:created_at;" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;" json:"updated_at"`

	Customer Customer `json:"customer" gorm:"foreignKey:CustomerId;references:ID"`

	// custom
	IsEmpty bool `gorm:"-" json:"-"`
}

func (t CustomerLimit) TableName() string {
	return "customer_limits"
}
