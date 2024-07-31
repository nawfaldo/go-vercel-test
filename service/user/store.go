package user

import (
	"database/sql"
	"vercer/types"
)

var db *sql.DB

func RegisterStore(database *sql.DB) {
	db = database
}

func GetUserByName(name string) *types.User {
	row := db.QueryRow("SELECT id FROM users WHERE name = $1", name)

	var user types.User
	row.Scan(&user.ID)

	return &user
}

func CreateUser(user types.User) error {
	_, err := db.Exec("INSERT INTO users (id, name, password) VALUES ($1, $2, $3)", user.ID, user.Name, user.Password)
	if err != nil {
		return err
	}

	return nil
}

func GetUserById(id string) any {
	row := db.QueryRow("SELECT id, name FROM users WHERE id = $1", id)

	var user types.User
	row.Scan(&user.ID, &user.Name)

	return &user
}
