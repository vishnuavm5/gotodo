package handlers

import (
	"encoding/json"
	"fmt"
	"hellocheck/internal/database"
	"hellocheck/pkg"
	"net/http"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Users struct {
	DB *database.Queries
}

func (user *Users) CreateNewUser(w http.ResponseWriter, r *http.Request) {

	user.DB = pkg.ConnectToDataBase()
	type parameters struct {
		Name     string `json:"name"`
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		pkg.RespondWithError(w, 400, fmt.Sprintf("Error parsing JSON:%s", err))
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error generating hashed password", err)
		return
	}
	fmt.Println(user.DB)
	id, err := user.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		Name:      params.Name,
		Username:  params.Username,
		Password:  string(hashedPassword),
		Email:     params.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		pkg.RespondWithError(w, 400, fmt.Sprintf("Couldn't create user:%s", err))
		return
	}

	pkg.RespondWithJSON(w, 201, id)

}

func (user *Users) Login(w http.ResponseWriter, r *http.Request) {

	user.DB = pkg.ConnectToDataBase()

	type parameters struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	type tokenData struct {
		Token string `json:"token"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		pkg.RespondWithError(w, 400, fmt.Sprintf("Error parsing JSON:%s", err))
		return
	}
	data, err := user.DB.GetUser(r.Context(), params.Username)

	if err != nil {
		pkg.RespondWithError(w, 400, fmt.Sprintf("No user found:%s", err))
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(params.Password))
	if err != nil {
		pkg.RespondWithError(w, 400, fmt.Sprintf("Invalid username or password: %s", err))
		return
	} else {
		token, err := pkg.GenerateJWT(params.Username, data.ID.String())
		if err != nil {
			pkg.RespondWithError(w, 400, fmt.Sprintf("Cannot login :%s", err))
			return
		}
		pkg.RespondWithJSON(w, 200, tokenData{Token: token})

	}

}
