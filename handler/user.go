package handler

import (
	"ca-dojo/model"
	"ca-dojo/util"
	"encoding/json"
	"fmt"
	"net/http"
)

func handleError (w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	// request body から name を取り出し userReq に格納
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	fmt.Println("body", body)

	var userReq model.UserCreateRequest
	json.Unmarshal(body, &userReq)

	// token 作成
	token, err := util.CreateToken(userReq.Name)
	if err != nil {
		handleError(w, err)
		return
	}
	fmt.Println("token", token)

	// user 構造体を定義して Create 処理
	user := model.User{Name: userReq.Name, Token: token}
	err = user.CreateUser()
	if err != nil {
		handleError(w, err)
		return
	}

	// response 返す
	resToken := model.UserCreateResponse{Token: token}
	res, err := json.Marshal(&resToken)
	if err != nil {
		handleError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func HandleGetUser(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("x-token")
	user, err := model.GetUser(token)
	if err != nil {
		handleError(w, err)
		return
	}

	userGet := model.UserGetResponse{Name: user.Name}
	res, err := json.Marshal(&userGet)
	if err != nil {
		handleError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func HandleUpdateUser(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("x-token")
	user, err := model.GetUser(token)
	if err != nil {
		handleError(w, err)
		return
	}

	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	var userReq model.UserUpdateRequest
	json.Unmarshal(body, &userReq)

	user.Name = userReq.Name
	if err := user.UpdateUser(); err != nil {
		handleError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
