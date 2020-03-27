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
func (user *User) SaveUser() error {
	tx, err := utils.DB.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()
	stmt, err := tx.Prepare("INSERT INTO user(username , email , password) VALUES ( ?,?,?)")

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(user.Username, user.Email, user.Password)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return err
}

func (user *User) FindUserByEmail(email string) error {

	err := utils.DB.Get(user, "SELECT * FROM user WHERE email = ? ", email)
	if err != nil {
		return err
	}
	return err
}
