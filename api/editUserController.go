package api

import (
	"crudinmemory/customerrs"
	"crudinmemory/repositories"
	"crudinmemory/services"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func editUserController(database map[uuid.UUID]repositories.User) http.HandlerFunc {
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
		idStr := chi.URLParam(r, "id")

		idUUID, err := uuid.Parse(idStr)

		if err != nil {
			sendJSON(w, Response{Data: customerrs.ErrInvalidUUID.Error(), Error: err.Error()}, http.StatusUnprocessableEntity)
			return
		}

		userRepository := repositories.NewUserInMemoryRepository(database)
		editUserService := services.NewEditUserService(userRepository)
		editedUser := editUserService.Execute(idUUID, body.FirstName, body.LastName, body.Biography)

		sendJSON(w, Response{Data: editedUser}, http.StatusAccepted)
	}
}
