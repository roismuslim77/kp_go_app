package entity

import (
	"time"
)

type Category struct {
	ID   int    `gorm:"column:id;type:int;primaryKey;autoIncrement:true;unique" json:"id"`
	Name string `gorm:"column:name;type:string;size:255;unique" json:"name"`

	CreatedAt time.Time `gorm:"column:created_at;" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;" json:"updated_at"`

	// custom
	IsEmpty bool `gorm:"-" json:"-"`
}

func (t Category) TableName() string {
	return "categories"
}
