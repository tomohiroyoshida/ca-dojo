package main

import (
	"ca-dojo/handler"
	"ca-dojo/model"
	"fmt"
	"net/http"
)

func main() {
	server := http.Server{
		Addr: ":8081",
	}
  model.Init()

	http.HandleFunc("/user/create", handler.HandleCreateUser)
	http.HandleFunc("/user/get", handler.HandleGetUser)
	http.HandleFunc("/user/update", handler.HandleUpdateUser)

	http.HandleFunc("/gacha/draw", handler.HandleDrawGacha)
	http.HandleFunc("/character/list", handler.HandleCharacterList)

	fmt.Println("Server listening on port 8081")
	server.ListenAndServe()
}
