package model

import (
	"fmt"
	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	DateOfBirth string    `json:"date_of_birth"`
}

func (u *User) ToString() string {
	return fmt.Sprintf("id=%d, name=%s, email=%s, dateOfBirth=%s", u.ID, u.Name, u.Email, u.DateOfBirth)
}
