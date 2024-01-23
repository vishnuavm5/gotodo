package pkg

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

func GenerateJWT(username string, id string) (string, error) {

	godotenv.Load()
	secret := os.Getenv("SECRET")
	println(secret)
	if secret == "" {
		log.Fatal("no secret found")

		return "", fmt.Errorf("no secret found")
	}

	//claims := data{Username: username, Exp: time.Now().Add(time.Hour * 1)}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"username": username, "id": id}).SignedString([]byte(secret))

	if err != nil {
		return "", fmt.Errorf("error in making token%s", err)
	}
	return token, nil

}
