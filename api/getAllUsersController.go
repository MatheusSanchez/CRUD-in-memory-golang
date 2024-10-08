package api

import (
	"crudinmemory/repositories"
	"crudinmemory/services"
	"net/http"

	"github.com/google/uuid"
)

func getAllUsersController(database map[uuid.UUID]repositories.User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		userRepository := repositories.NewUserInMemoryRepository(database)
		getAllUsersService := services.NewGetAllUsersService(userRepository)

		users := getAllUsersService.Execute()

		sendJSON(w, Response{Data: users}, http.StatusAccepted)
	}
}
