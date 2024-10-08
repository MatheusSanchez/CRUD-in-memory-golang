package api

import (
	"crudinmemory/customerrs"
	"crudinmemory/repositories"
	"crudinmemory/services"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func deleteUserController(database map[uuid.UUID]repositories.User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		idStr := chi.URLParam(r, "id")

		idUUID, err := uuid.Parse(idStr)

		if err != nil {
			sendJSON(w, Response{Data: customerrs.ErrInvalidUUID.Error(), Error: err.Error()}, http.StatusUnprocessableEntity)
			return
		}

		userRepository := repositories.NewUserInMemoryRepository(database)
		deleteUserService := services.NewDeleteUserService(userRepository)

		deletedUser, err := deleteUserService.Execute(idUUID)

		if err != nil {

			if errors.Is(err, customerrs.ErrUserNotFoundById) {
				sendJSON(w, Response{Error: err.Error()}, http.StatusNotFound)
				return
			}
			sendJSON(w, Response{Error: customerrs.ErrSomethingWentWrong.Error()}, http.StatusInternalServerError)
		}

		sendJSON(w, Response{Data: deletedUser}, http.StatusAccepted)
	}
}
