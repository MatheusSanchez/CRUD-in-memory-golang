package api

import (
	"crudinmemory/repositories"
	"crudinmemory/services"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/google/uuid"
)

type CreateUserBodySchema struct {
	FirstName string `json:"firstname" validate:"required"`
	LastName  string `json:"lastname" validate:"required"`
	Biography string `json:"bio" validate:"required"`
}

func postUserController(database map[uuid.UUID]repositories.User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var body CreateUserBodySchema

		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			sendJSON(w, Response{Error: "invalid body"}, http.StatusUnprocessableEntity)
			return
		}

		if err := validate.Struct(body); err != nil {
			sendJSON(w, Response{Error: "Please provide FirstName LastName and bio for the user"}, http.StatusBadRequest)
			return
		}

		userRepository := repositories.NewUserInMemoryRepository(database)
		createUserService := services.NewCreateUserService(userRepository)
		newUser := createUserService.Execute(body.FirstName, body.LastName, body.Biography)

		slog.Info("new user was created", "userId", newUser.Id)
		sendJSON(w, Response{Data: newUser}, http.StatusAccepted)
	}
}
