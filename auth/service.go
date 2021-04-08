package auth

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

type Service interface {
	GenerateToken(NipGuru string) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {
}

func NewService() *jwtService {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &jwtService{}
}

func (s *jwtService) GenerateToken(NipGuru string) (string, error) {
	claim := jwt.MapClaims{}

	claim["guru_nip"] = NipGuru
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		fmt.Println(err)
		return signedToken, err
	}

	return signedToken, nil
}

func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Invalid token")
		}

		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}
