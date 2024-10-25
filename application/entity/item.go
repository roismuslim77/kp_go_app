package entity

import (
	"time"
)

type Item struct {
	ID          int      `gorm:"column:id;type:int;primaryKey;autoIncrement:true;unique" json:"id"`
	CategoryId  int      `gorm:"column:category_id;type:int" json:"category_id"`
	Name        string   `gorm:"column:name;type:string;size:255" json:"name"`
	Description string   `gorm:"column:description;type:string;size:255" json:"description"`
	Price       *float64 `gorm:"column:price;type:float;" json:"price"`

	CreatedAt time.Time `gorm:"column:created_at;" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;" json:"updated_at"`

	// custom
	IsEmpty bool `gorm:"-" json:"-"`
}

func (t Item) TableName() string {
	return "items"
}
