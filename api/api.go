package api

import (
	"crudinmemory/repositories"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
)

var db = make(map[uuid.UUID]repositories.User)

func NewHandler() http.Handler {
	router := chi.NewMux()

	router.Use(middleware.Recoverer)
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)

	router.Use(jsonMiddleware)

	router.Get("/healthcheck", handleHealthCheck)

	router.Post("/users", postUserController(db))
	router.Put("/users/{id}", editUserController(db))
	router.Get("/users", getAllUsersController(db))
	router.Get("/user/{id}", getUserController(db))
	router.Delete("/user/{id}", deleteUserController(db))

	return router
}

var validate = validator.New()

func handleHealthCheck(w http.ResponseWriter, r *http.Request) {
	sendJSON(w, Response{Data: "ok"}, http.StatusAccepted)
}
