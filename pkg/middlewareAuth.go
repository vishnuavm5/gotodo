package pkg

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

type Claims struct {
	Id       uuid.UUID
	Username string
	jwt.StandardClaims
}

type UserData struct {
	Id       uuid.UUID
	Username string
}
type authedHandler func(http.ResponseWriter, *http.Request, UserData)

func MiddlewareAuth(handler authedHandler) http.HandlerFunc {
	claims := &Claims{}
	godotenv.Load()
	secret := os.Getenv("SECRET")
	println(secret)
	if secret == "" {
		log.Fatal("no secret found")
	}
	return func(w http.ResponseWriter, r *http.Request) {
		token, err := getJWTToken(r.Header)
		if err != nil {
			RespondWithError(w, 403, fmt.Sprintf("Auth error:%v", err))
			return
		}
		_, err = jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) { return []byte(secret), nil })

		if err != nil {
			RespondWithError(w, 400, fmt.Sprintf("Could't get user : %v ", err))
			return
		}
		handler(w, r, UserData{Id: claims.Id, Username: claims.Username})
	}
}
