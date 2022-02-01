package http

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/talkanbaev-artur/interview/config"
	"github.com/talkanbaev-artur/interview/user/service"
	"go.uber.org/zap"
)

func NewHTTPServer(cancel context.CancelFunc, config config.AppConfig, us service.Service) {
	r := mux.NewRouter()

	r.HandleFunc("/users", createHandleListUsers(us)).Methods("GET")
	r.HandleFunc("/users/{id}", createHandlerGetUser(us)).Methods("GET")
	r.HandleFunc("/users", createHandlerCreateUser(us)).Methods("POST")
	r.HandleFunc("/users/{id}", createHandlerUpdateUser(us)).Methods("PUT")
	r.HandleFunc("/users/{id}", createHandlerDeleteUser(us)).Methods("DELETE")

	addr := fmt.Sprintf("0.0.0.0:%d", config.ServerPort)
	server := http.Server{
		Addr:    addr,
		Handler: r,
	}

	zap.S().Infow("Server is running on: " + addr)
	zap.S().Errorw("got error", "error", server.ListenAndServe())
	cancel()
}
