package auth

import (
	"food-truck-api/package/auth/contract"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type Repository interface {
	HashPassword(password string) (string, error)
	CheckPassword(hash string, password string) bool
	CreateToken(payload *contract.CreateTokenRequest) (string, error)
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

func (r *repository) CheckPassword(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (r *repository) CreateToken(payload *contract.CreateTokenRequest) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = payload.Name
	claims["email"] = payload.Email
	claims["id"] = payload.ID

	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	return token.SignedString([]byte(os.Getenv("SECRET")))
}
