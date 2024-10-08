package api

import (
	"crudinmemory/customerrs"
	"encoding/json"
	"log/slog"
	"net/http"
)

type Response struct {
	Error string `json:"error,omitempty"`
	Data  any    `json:"data,omitempty"`
}

func jsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func sendJSON(w http.ResponseWriter, response Response, status int) {

	data, err := json.Marshal(response)

	if err != nil {
		slog.Error("err when trying to marshal response", "error", err)
		sendJSON(w, Response{Error: customerrs.ErrSomethingWentWrong.Error()}, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)

	_, err = w.Write(data)

	if err != nil {
		slog.Error("err when trying to send response", "error", err)
		panic(err)
	}

}
