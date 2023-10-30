package operations

import (
	"api/db"
	"api/models"
)

func CreateUser(user models.User) (id int64, err error) {
	connection, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer connection.Close()

	SQL := `INSERT INTO users (name, age, gender) VALUES ($1, $2, $3) RETURNING id`
	err = connection.QueryRow(SQL, user.Name, user.Age, user.Gender).Scan(&id)

	return
}
