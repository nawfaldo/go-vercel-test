package user

import "database/sql"

func GetUsers(database *sql.DB) any {
	rows, _ := db.Query("SELECT name FROM users")
	defer rows.Close()

	type User struct {
		Name string `json:"name"`
	}

	var users []User

	for rows.Next() {
		var u User

		rows.Scan(&u.Name)

		users = append(users, u)
	}

	return users
}
