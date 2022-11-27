package user

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"user-service/internal/user-service/apperror"
	"user-service/internal/user-service/handlers"
	"user-service/pkg/logging"
)

const (
	usersURL = "/users"
	userURL  = "/users/:uuid"
)

type handler struct {
	logger *logging.Logger
}

func NewHandler(logger *logging.Logger) handlers.Handler {
	return &handler{
		logger: logger,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, usersURL, apperror.Middleware(h.GetList))
	router.HandlerFunc(http.MethodPost, usersURL, apperror.Middleware(h.CreateUser))
	router.HandlerFunc(http.MethodGet, userURL, apperror.Middleware(h.GetUserByUUID))
	router.HandlerFunc(http.MethodPut, userURL, apperror.Middleware(h.UpdateUser))
	router.HandlerFunc(http.MethodPatch, userURL, apperror.Middleware(h.PartialUpdateUser))
	router.HandlerFunc(http.MethodDelete, userURL, apperror.Middleware(h.DeleteUser))
}

func (h *handler) GetList(writer http.ResponseWriter, request *http.Request) error {
	writer.WriteHeader(200)
	writer.Write([]byte("this is list of users"))

	return nil
}

func (h *handler) CreateUser(writer http.ResponseWriter, request *http.Request) error {
	writer.WriteHeader(201)
	writer.Write([]byte("this is create user"))

	return nil
}

func (h *handler) GetUserByUUID(writer http.ResponseWriter, request *http.Request) error {
	writer.WriteHeader(200)
	writer.Write([]byte("this is user by uuid"))

	return nil
}

func (h *handler) UpdateUser(writer http.ResponseWriter, request *http.Request) error {
	writer.WriteHeader(204)
	writer.Write([]byte("this is update user"))

	return nil
}

func (h *handler) PartialUpdateUser(writer http.ResponseWriter, request *http.Request) error {
	writer.WriteHeader(204)
	writer.Write([]byte("this is partial update user"))

	return nil
}

func (h *handler) DeleteUser(writer http.ResponseWriter, request *http.Request) error {
	writer.WriteHeader(204)
	writer.Write([]byte("this is delete user"))

	return nil
}
