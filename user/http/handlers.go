package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/talkanbaev-artur/interview/user/service"
	"go.uber.org/zap"
)

func createHandleListUsers(us service.Service) http.HandlerFunc {
	logr := zap.S().With("level", "httpServer")
	return func(rw http.ResponseWriter, r *http.Request) {
		users, err := us.ListUser(r.Context())
		if err != nil {
			logr.Errorw("error happended during fetch process", "error", err.Error())
			http.Error(rw, fmt.Sprintf("error happended during fetch process, erorr: %s", err.Error()), 500)
			return
		}
		rw.WriteHeader(200)
		json.NewEncoder(rw).Encode(&users)
		logr.Infow("Succesfully returned users list answer")
	}
}

func createHandlerGetUser(us service.Service) http.HandlerFunc {
	logr := zap.S().With("level", "httpServer")
	return func(rw http.ResponseWriter, r *http.Request) {

		params := mux.Vars(r)
		id, _ := strconv.ParseInt(params["id"], 10, 64)

		users, err := us.GetUserByID(r.Context(), id)
		if err != nil {
			logr.Errorw("error happended during fetch process", "error", err.Error())
			http.Error(rw, fmt.Sprintf("error happended during fetch process, erorr: %s", err.Error()), 500)
			return
		}

		rw.WriteHeader(200)
		json.NewEncoder(rw).Encode(&users)
		logr.Infow("Succesfully returned users list answer")
	}
}

func createHandlerCreateUser(us service.Service) http.HandlerFunc {
	logr := zap.S().With("level", "httpServer")
	return func(rw http.ResponseWriter, r *http.Request) {
		var req service.UserChangeInput
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			logr.Errorw("failed to parse request params", "error", err.Error())
			http.Error(rw, fmt.Sprintf("failed to parse request params, erorr: %s", err.Error()), 500)
			return
		}

		user, err := us.RegisterUser(r.Context(), req)
		if err != nil {
			logr.Errorw("error happended during fetch process", "error", err.Error())
			http.Error(rw, fmt.Sprintf("error happended during fetch process, erorr: %s", err.Error()), 500)
			return
		}

		rw.WriteHeader(201)
		json.NewEncoder(rw).Encode(&user)
		logr.Infow("Succesfully returned users list answer")
	}
}

func createHandlerUpdateUser(us service.Service) http.HandlerFunc {
	logr := zap.S().With("level", "httpServer")
	return func(rw http.ResponseWriter, r *http.Request) {
		var req service.UserChangeInput
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			logr.Errorw("failed to parse request params", "error", err.Error())
			http.Error(rw, fmt.Sprintf("failed to parse request params, erorr: %s", err.Error()), 500)
			return
		}

		params := mux.Vars(r)
		id, _ := strconv.ParseInt(params["id"], 10, 64)

		err = us.UpdateUserAccount(r.Context(), id, req)
		if err != nil {
			logr.Errorw("error happended during fetch process", "error", err.Error())
			http.Error(rw, fmt.Sprintf("error happended during fetch process, erorr: %s", err.Error()), 500)
			return
		}

		rw.WriteHeader(200)
		logr.Infow("Succesfully returned users list answer")
	}
}

func createHandlerDeleteUser(us service.Service) http.HandlerFunc {
	logr := zap.S().With("level", "httpServer")
	return func(rw http.ResponseWriter, r *http.Request) {

		params := mux.Vars(r)
		id, _ := strconv.ParseInt(params["id"], 10, 64)

		err := us.SuspendUser(r.Context(), id)
		if err != nil {
			logr.Errorw("error happended during fetch process", "error", err.Error())
			http.Error(rw, fmt.Sprintf("error happended during fetch process, erorr: %s", err.Error()), 500)
			return
		}

		rw.WriteHeader(204)
		logr.Infow("Succesfully returned users list answer")
	}
}
