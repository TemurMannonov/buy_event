package models

type Log struct {
	ID        string `json:"id" db:"id"`
	Message   string `json:"message" db:"message"`
	CreatedAt string `json:"created_at" db:"created_at"`
}
