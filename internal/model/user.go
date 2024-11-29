package model

import "github.com/google/uuid"

type User struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	DateOfBirth string    `json:"date_of_birth"`
}
