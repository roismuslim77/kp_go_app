package auth

import (
	"time"
)

type RegisterUserRequest struct {
	NIK               string    `json:"nik" binding:"required"`
	FullName          string    `json:"full_name" binding:"required"`
	LegalName         string    `json:"legal_name" binding:"required"`
	BirthPlace        string    `json:"birth_place" binding:"required"`
	BirthDate         time.Time `json:"birth_date" binding:"required"`
	Salary            float64   `json:"salary" binding:"required"`
	IdentityCardLink  string    `json:"identity_card_link" binding:"required"`
	IdentityPhotoLink string    `json:"identity_photo_link" binding:"required"`
	Password          string    `json:"password" binding:"required"`
}

type LoginCustomerReq struct {
	NIK      string `json:"nik" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CustomerAuthHeaderReq struct {
	Authorization string `header:"Authorization" binding:"required"`
}
