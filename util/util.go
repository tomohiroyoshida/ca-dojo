package util

import (
	"ca-dojo/model"
	"log"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

func CreateToken(name string) (string, error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		log.Println(err)
		return "", nil
	}
  token := uuid.String()
	return token, nil
}

func WeightPick() (gachaResult model.GachaResult, err error) {
  //characterを全部取得(重み昇順でソート)
  characters, err := getAllCharacters()
  if err != nil {
    return
  }

  //重みを合計する
  total_weight, err := sumWeight(characters)
  if err != nil {
    return
  }

  //乱数生成
  rand.Seed(time.Now().UnixNano())
  rnd := rand.Intn(int(total_weight))

  //生成された数字に基づいて返すcharacterを決める
  var picked model.CharacterWithWeight
  for i:=0; i<len(characters); i+=1 {
    if rnd < int(characters[i].Weight) {
      picked = characters[i]
      break
    }
    rnd -= int(characters[i].Weight)
  }

  gachaResult = model.GachaResult{CharacterID: picked.ID, Name: picked.Name}
  return
}

func getAllCharacters() (characters []model.CharacterWithWeight, err error) {
  rows, err := model.Db.Query("select id, name, weight from characters order by weight ASC")
  for rows.Next() {
    character := model.CharacterWithWeight{}
    err = rows.Scan(&character.ID, &character.Name, &character.Weight)
    if err != nil {
      return
    }
    characters = append(characters, character)
  }
  return
}

func sumWeight(characters []model.CharacterWithWeight) (totalWeight int, err error) {
  for i:=0; i<len(characters); i+=1 {
    totalWeight += characters[i].Weight
  }
  return
}