package user

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"user-service/internal/user-service/handlers"
)

const (
	usersURL = "/users"
	userURL  = "/users/:uuid"
)

type handler struct {
}

func NewHandler() handlers.Handler {
	return &handler{}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(usersURL, h.GetList)
	router.POST(userURL, h.GetUserByUUID)
	router.GET(usersURL, h.CreateUser)
	router.PUT(userURL, h.UpdateUser)
	router.PATCH(userURL, h.PartialUpdateUser)
	router.DELETE(userURL, h.DeleteUser)
}

func (h *handler) GetList(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.Write([]byte("this is list of users"))
}

func (h *handler) GetUserByUUID(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.Write([]byte("this is user by uuid"))
}

func (h *handler) CreateUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.Write([]byte("this is create user"))
}

func (h *handler) UpdateUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.Write([]byte("this is update user"))
}

func (h *handler) PartialUpdateUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.Write([]byte("this is partial update user"))
}

func (h *handler) DeleteUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.Write([]byte("this is delete user"))
}
