package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net"
	"net/http"
	"time"
	"user-service/internal/user-service/user"
)

func main() {
	fmt.Println("create router")
	router := httprouter.New()

	handler := user.NewHandler()
	handler.Register(router)

	start(router)
}

func start(router *httprouter.Router) {

	listener, err := net.Listen("tcp", "127.0.0.1:1234")
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	log.Fatal(server.Serve(listener))
}
