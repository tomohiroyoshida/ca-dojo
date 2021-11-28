package model

import "time"

// Userの構造体
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated__at"`
}

type UserCreateRequest struct {
	Name string `json:"name"`
}

type UserCreateResponse struct {
	Token string `json:"token"`
}

type UserGetResponse struct {
	Name string `json:"name"`
}

type UserUpdateRequest struct {
	Name string `json:"name"`
}

// ガチャ
type GachaResult struct {
	CharacterID int    `json:"characterID"`
	Name        string `json:"name"`
}

type GachaDrawRequest struct {
	Times int `json:"times"`
}

type GachaDrawResponse struct {
	Results []GachaResult `json:"results"`
}

// キャラ
type UserCharacter struct {
	UserCharacterID int    `json:"userCharacterID"`
	CharacterID     int    `json:"characterID"`
	Name            string `json:"name"`
}

type CharacterListResponse struct {
	Characters []UserCharacter `json:"characters"`
}
type Character struct {
	ID     int
	Name   string
	Weight int
}

type UserCharacterWithUserID struct {
	ID          int `json:"userCharacterID"`
	CharacterID int `json:"characterID"`
	UserID      int `json:"userID"`
}

type UserCharacterListResponse struct {
	UserCharacterID int    `json:"userCharacterID"`
	CharacterID     int    `json:"characterID"`
	Name            string `json:"name"`
}
