package models

import (
	"JWT-Based-Authentication/utils"
)

type User struct {
	Id       uint
	Username string
	Email    string
	Password string
}

func CreateTable() {
	_, err := utils.DB.Exec(`CREATE TABLE user (id integer NOT NULL  UNIQUE AUTO_INCREMENT, username varchar(255) UNIQUE ,email varchar(255) UNIQUE , password text,PRIMARY KEY (id));`)

	if err != nil {
		panic(err)
	}
}
func (user *User) SaveUser() {
	utils.DB.NamedExec("INSERT INTO user VALUES ( :id, :username,:email, :password)", user)
}

func (user *User) FindUserByEmail(email string) {
	utils.DB.Get(user, "SELECT * FROM user WHERE email = ? ", email)
}
