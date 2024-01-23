package handlers

import (
	"encoding/json"
	"fmt"
	"hellocheck/internal/database"
	"hellocheck/models"
	"hellocheck/pkg"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type Todo struct {
	DB *database.Queries
}

func (todo *Todo) CreateTodo(w http.ResponseWriter, r *http.Request, user pkg.UserData) {
	type parameters struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	decoder := json.NewDecoder(r.Body)
	params := &parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		pkg.RespondWithError(w, 400, fmt.Sprintf("Error parsing JSON:%s", err))
		return
	}
	todo.DB = pkg.ConnectToDataBase()
	todoData, err := todo.DB.CreateTodo(r.Context(), database.CreateTodoParams{Title: params.Title, Description: params.Description, UserID: user.Id, CreatedAt: time.Now(), UpdatedAt: time.Now(), ID: uuid.New()})
	if err != nil {
		pkg.RespondWithError(w, 403, fmt.Sprintf("Error adding todo:%s", err))
		return
	}

	pkg.RespondWithJSON(w, 201, todoData)

}
func (todo *Todo) GetAllTodos(w http.ResponseWriter, r *http.Request, user pkg.UserData) {
	todo.DB = pkg.ConnectToDataBase()
	todos, err := todo.DB.GetTodoList(r.Context(), user.Id)

	if err != nil {
		pkg.RespondWithError(w, 403, fmt.Sprintf("something went wrong couldnt fetch todos,%s", err))
		return
	}
	pkg.RespondWithJSON(w, 200, models.DatabaseListToList(todos))

}

func (todo *Todo) GetTodoById(w http.ResponseWriter, r *http.Request, user pkg.UserData) {
	todo.DB = pkg.ConnectToDataBase()
	// type parameters struct {
	// 	Id uuid.UUID `json:"id"`
	// }
	// decoder := json.NewDecoder(r.Body)
	// params := &parameters{}
	// err := decoder.Decode(&params)

	Id := chi.URLParam(r, "id")
	TodoId, err := uuid.Parse(Id)
	if err != nil {
		pkg.RespondWithError(w, 400, fmt.Sprintf("Error parsing JSON:%s", err))
		return

	}
	dbTodo, err := todo.DB.GetTodoById(r.Context(), TodoId)
	if err != nil {
		pkg.RespondWithError(w, 403, fmt.Sprintf("Unable to get todos:%s", err))
	}
	pkg.RespondWithJSON(w, 200, models.DatabaseTodoToTodo(dbTodo))
}

func (todo *Todo) UpdateTodoById(w http.ResponseWriter, r *http.Request, user pkg.UserData) {
	type parameters struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}
	todo.DB = pkg.ConnectToDataBase()
	todoId, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		pkg.RespondWithError(w, 400, fmt.Sprintf("Error parsing UUID:%s", err))
		return
	}

	decoder := json.NewDecoder(r.Body)
	params := &parameters{}
	err = decoder.Decode(&params)
	if err != nil {
		pkg.RespondWithError(w, 400, fmt.Sprintf("Error parsing json:%s", err))
		return
	}

	dbTodo, err := todo.DB.UpdateTodoById(r.Context(), database.UpdateTodoByIdParams{ID: todoId, Title: params.Title, Description: params.Description, UserID: user.Id})
	if err != nil {
		pkg.RespondWithError(w, 400, fmt.Sprintf("Error%s", err))
		return

	}
	pkg.RespondWithJSON(w, 201, dbTodo)

}

func (todo *Todo) DeleteById(w http.ResponseWriter, r *http.Request, user pkg.UserData) {
	todo.DB = pkg.ConnectToDataBase()
	type Response struct {
		Message string `json:"message"`
	}
	TodoId, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		pkg.RespondWithError(w, 400, fmt.Sprintf("error parsing uuid %s", err))
	}
	err = todo.DB.DeleteTodoById(r.Context(), database.DeleteTodoByIdParams{ID: TodoId, UserID: user.Id})
	if err != nil {
		pkg.RespondWithError(w, 405, fmt.Sprintf("Error deleting todo : %s", err))
		return
	}

	pkg.RespondWithJSON(w, 201, Response{Message: "success"})

}
