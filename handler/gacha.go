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
	fmt.Println("user", user)

	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)

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
		
		userCharacter := model.UserCharacterWithUserID{
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

func HandleCharacterList (w http.ResponseWriter, r *http.Request) () {
	token := r.Header.Get("x-token")
	user, err := model.GetUser(token)
	if err != nil {
		handleError(w, err)
		return
	}
	fmt.Println("user", user)

	userCharacterListResponses, err := util.GetAllUserCharacters(user)
	if err != nil {
		handleError(w, err)
		return
	}

	characterListResponse := model.CharacterListResponse{
    Characters: userCharacterListResponses,
  }

	data, err := json.Marshal(&characterListResponse)
	if err != nil {
		handleError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
  w.Write(data)
}