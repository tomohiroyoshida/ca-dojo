package handler

import (
	"ca-dojo/model"
	"ca-dojo/util"
	"encoding/json"
	"fmt"
	"net/http"
)

func HandleDrawGacha (w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("x-token")
	user, err := model.GetUser(token)
	if err != nil {
		handleError(w, err)
		return
	}

	// request body から name を取り出し userReq に格納
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	fmt.Println("body", body)

	var gachaReq model.GachaDrawRequest
	json.Unmarshal(body, &gachaReq)
	times := gachaReq.Times

	var results []model.GachaResult

	for ; times > 0; times -= 1 {
		gachaResult, err := util.WeightPick()
		if err != nil {
			handleError(w, err)
			return
		}

		userCharacter := model.UserCharacter{
			UserID: user.ID,
			CharacterID: gachaResult.CharacterID,
    }
		userCharacter.CreateUserCharacter()
		results = append(results, gachaResult)
	} 

	gachaDrawResponse := model.GachaDrawResponse{
    Results: results,
  }
	output, err := json.Marshal(&gachaDrawResponse)
  if err != nil {
		handleError(w, err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
  w.Write(output)
}