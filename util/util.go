package util

import (
	"ca-dojo/model"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

func handleError (err error) {
	log.Fatal(err)
}

func CreateToken(name string) (string, error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		handleError(err)
		return "", nil
	}
  token := uuid.String()
	return token, nil
}

func WeightPick() (gachaResult model.GachaResult, err error) {
  //characterを重み(weight)昇順で全部取得
  characters, err := getAllCharacters()
  if err != nil {
    handleError(err)
    return
  }

  //重みを合計する
  totalWeight, err := sumWeight(characters)
  if err != nil {
    handleError(err)
    return
  }

  //乱数生成
  rand.Seed(time.Now().UnixNano())
  randInt := rand.Intn(int(totalWeight))
  fmt.Println("int", randInt)

  //生成された数字に基づいて返すcharacterを決める
  var resCharacter model.Character
  for i:=0; i<len(characters); i+=1 {
    if randInt > characters[i].Weight {
      resCharacter = characters[i]
      break
    }
    randInt += characters[i].Weight
  }

  gachaResult = model.GachaResult{CharacterID: resCharacter.ID, Name: resCharacter.Name}
  return
}

func getAllCharacters() (characters []model.Character, err error) {
  rows, err := model.Db.Query("select id, name, weight from characters order by weight ASC")
  for rows.Next() {
    character := model.Character{}
    err = rows.Scan(&character.ID, &character.Name, &character.Weight)
    if err != nil {
      handleError(err)
      return
    }
    characters = append(characters, character)
  }
  return
}

func sumWeight(characters []model.Character) (totalWeight int, err error) {
  for i:=0; i<len(characters); i+=1 {
    totalWeight += characters[i].Weight
  }
  return
}

func GetAllUserCharacters (user model.User) (userCharacterListResponses []model.UserCharacter, err error) {
  rows, err := model.Db.Query("select id, character_id from user_characters where user_id = ?", user.ID)
  if err != nil {
    handleError(err)
    return
  }
  for rows.Next() {
    characterListRes := model.UserCharacter{}
    err = rows.Scan(&characterListRes.UserCharacterID, &characterListRes.CharacterID)
    if err != nil {
      handleError(err)
      return
    }

    var character model.Character
    character, err = getCharacter(characterListRes.CharacterID)
    if err != nil {
      handleError(err)
      return
    }
  
    characterListRes.Name = character.Name
    userCharacterListResponses = append(userCharacterListResponses, characterListRes)
    fmt.Println("character response: ", characterListRes)
  }
  return 
}

func getCharacter(characterID int) (character model.Character, err error) {
  err = model.Db.QueryRow("select name from characters where id = ?", characterID).Scan(&character.ID)
  if err != nil {
      handleError(err)
      return
  }
  return
}
