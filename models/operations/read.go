package operations

import (
	"api/db"
	"api/models"
)

func GetUserByID(id int64) (user models.User, err error) {
	connection, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer connection.Close()

	SQL := `SELECT * FROM users WHERE id=$1`
	row := connection.QueryRow(SQL, id)
	err = row.Scan(&user.ID, &user.Name, &user.Age, &user.Gender)

	return
}
