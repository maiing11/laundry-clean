package entities

import "github.com/google/uuid"

type TblCustomer struct {
	ID          uuid.UUID `json:"id,omitempty"`
	Name        string    `json:"name" form:"name"`
	Address     string    `json:"address" form:"name"`
	PhoneNumber string    `json:"phoneNumber" form:"phoneNumber"`
}
