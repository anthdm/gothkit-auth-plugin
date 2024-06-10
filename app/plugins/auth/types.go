package auth

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID              int `bun:"id,pk,autoincrement"`
	Email           string
	FirstName       string
	LastName        string
	PasswordHash    string
	EmailVerifiedAt time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func createUserFromFormValues(values SignupFormValues) (User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(values.Password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, err
	}
	return User{
		Email:        values.Email,
		FirstName:    values.FirstName,
		LastName:     values.LastName,
		PasswordHash: string(hash),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}, nil
}
