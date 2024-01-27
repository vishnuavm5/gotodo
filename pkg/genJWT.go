package pkg

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-jwt/jwt"
)

func GenerateJWT(username string, id string) (string, error) {

	secret := os.Getenv("SECRET")

	if secret == "" {
		log.Fatal("no secret found")

		return "", fmt.Errorf("no secret found")
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"username": username, "id": id}).SignedString([]byte(secret))

	if err != nil {
		return "", fmt.Errorf("error in making token%s", err)
	}
	return token, nil

}
