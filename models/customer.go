package models

type Customer struct {
	ID          string `json:"id" db:"id"`
	PhoneNumber string `json:"phone_number" db:"phone_number"`
	Email       string `json:"email" db:"email"`
}
