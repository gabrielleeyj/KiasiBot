package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	secretKey []byte
)

type User struct {
	ChatID    string   `json:"chatID"`
	Locations Location `json:"locations"`
}

type Location struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type Data struct {
	ChatID    string `json:"chatID"`
	Locations Location
}
type Claims struct {
	Data Data `json:"data,omitempty"`
	jwt.StandardClaims
}

func init() {
	secretKey = []byte(os.Getenv("SECRET_KEY"))
}

func Encode(body Data) (string, int64, error) {
	expiry := time.Now().Add(time.Hour * 24).Unix()
	claims := Claims{
		Data{
			ChatID:    body.ChatID,
			Locations: body.Locations,
		},
		jwt.StandardClaims{
			ExpiresAt: expiry,
			Issuer:    "some.issuer.com",
			Subject:   body.ChatID,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	result, err := token.SignedString(secretKey)
	if err != nil {
		fmt.Println(err)
	}
	return result, expiry, nil
}
