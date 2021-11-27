package model

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

// DBへの接続
func Init() {
	var err error
	Db, err = sql.Open("mysql", "root:root@tcp(mysql:3306)/ca_dojo?parseTime=true")
	if err != nil {
		panic(err)
	}
	log.Println("DB conndected!")
}

func handleError (err error) {
	log.Fatal(err)
}

// ユーザー
func GetUser(token string) (user User, err error) {
	user = User{}
	err = Db.QueryRow("select id, name, token, created_at, updated_at from users where token = ?", token).
	Scan(&user.ID, &user.Name, &user.Token, &user.CreatedAt, &user.UpdatedAt)
	fmt.Println("user", user)
	if err != nil {
		handleError((err))
		return
	}
	return
}

func (user *User) CreateUser() (err error) {
	if _, err = Db.Exec("insert into users (name, token) values (?, ?)", user.Name, user.Token); err != nil {
		handleError((err))
		return
	}
	return
}

func (user *User) UpdateUser() (err error) {
	if _, err = Db.Exec("update users set name = ? where id = ?", user.Name, user.ID); err != nil {
		handleError((err))
		return
	}
	return
}

// ガチャ
func (userCharacter *UserCharacterWithUserID) CreateUserCharacter() (err error) {
	_, err = Db.Exec("insert into user_characters (user_id, character_id) values (?, ?)", 
									  userCharacter.UserID, userCharacter.CharacterID)
  return
}
