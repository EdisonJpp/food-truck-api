package auth

import (
	"golang.org/x/crypto/bcrypt"
)

type Repository interface {
	HashPassword(password string) (string, error)
}

type repository struct {
}

func NewRepo() Repository {
	return &repository{}
}

func (r *repository) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
