package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/talkanbaev-artur/interview/user/service"
	"go.uber.org/zap"
)

type errorMsg struct {
	Message string `json:"msg"`
	Details string `json:"err_details"`
}

func answerErr(err errorMsg, rw http.ResponseWriter, logr *zap.SugaredLogger) {
	rw.WriteHeader(500)
	logr.Errorw(err.Message, "error", err.Details)
	json.NewEncoder(rw).Encode(err)
}

func createHandleListUsers(us service.Service) http.HandlerFunc {
	logr := zap.S().With("level", "httpServer")
	return func(rw http.ResponseWriter, r *http.Request) {
		users, err := us.ListUser(r.Context())
		if err != nil {
			answerErr(errorMsg{"error happended during fetch process", err.Error()}, rw, logr)
			return
		}
		rw.WriteHeader(200)
		json.NewEncoder(rw).Encode(users)
		logr.Infow("Succesfully returned users list answer")
	}
}

func createHandlerGetUser(us service.Service) http.HandlerFunc {
	logr := zap.S().With("level", "httpServer")
	return func(rw http.ResponseWriter, r *http.Request) {

		params := mux.Vars(r)
		id, _ := strconv.ParseInt(params["id"], 10, 64)

		users, err := us.GetUserByID(r.Context(), id)
		if err == service.ErrUserNotFound {
			rw.WriteHeader(204)
			return
		}
		if err != nil {
			answerErr(errorMsg{"error happended during fetch process", err.Error()}, rw, logr)
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
			answerErr(errorMsg{"failed to parse request params", err.Error()}, rw, logr)
			return
		}

		user, err := us.RegisterUser(r.Context(), req)
		if err != nil {
			answerErr(errorMsg{"error happended during registration process", err.Error()}, rw, logr)
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
			answerErr(errorMsg{"failed to parse request params", err.Error()}, rw, logr)
			return
		}

		params := mux.Vars(r)
		id, _ := strconv.ParseInt(params["id"], 10, 64)

		err = us.UpdateUserAccount(r.Context(), id, req)
		if err != nil {
			answerErr(errorMsg{"error happended during update process", err.Error()}, rw, logr)
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
			answerErr(errorMsg{"error happended during delete process", err.Error()}, rw, logr)
			return
		}

		rw.WriteHeader(204)
		logr.Infow("Succesfully returned users list answer")
	}
}
