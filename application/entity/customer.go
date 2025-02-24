package entity

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type Customer struct {
	ID                int       `gorm:"column:id;type:int;primaryKey;autoIncrement:true;unique" json:"id"`
	Nik               string    `gorm:"column:nik;type:string;size:255;unique" json:"nik"`
	FullName          string    `gorm:"column:name;type:string;size:255" json:"name"`
	LegalName         string    `gorm:"column:legal_name;type:string;size:255" json:"legal_name"`
	BirthPlace        string    `gorm:"column:birth_place;type:string;size:255" json:"birth_place"`
	BirthDate         time.Time `gorm:"column:birth_date;" json:"birth_date"`
	Salary            float64   `gorm:"column:salary;type:float;" json:"salary"`
	IdentityCardLink  string    `gorm:"column:identity_card_link;type:string;size:255" json:"identity_card_link"`
	IdentityPhotoLink string    `gorm:"column:identity_photo_link;type:string;size:255" json:"identity_photo_link"`

	CreatedAt time.Time `gorm:"column:created_at;" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;" json:"updated_at"`

	// custom
	IsEmpty bool `gorm:"-" json:"-"`
}

func (t Customer) TableName() string {
	return "customers"
}

type Claims struct {
	CustomerId int `json:"customer_id"`
	jwt.RegisteredClaims
}
