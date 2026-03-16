package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID			uuid.UUID	`json:"id"`
	Email 		string		`json:"email"`
	Password	string		`json:"password"`
	Username    *string		`json:"username"`
	Phone		*int		`json:"phone"`
	Location 	*string		`json:"location"`
	Role		*string		`json:"role"`
	CreatedAt	time.Time	`json:"created_at"`
	UpdatedAt	time.Time	`json:"updated_at"`
}