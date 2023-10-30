package operations

import (
	"api/db"
	"api/models"
)

func UpdateUserByID(id int64, user models.User) (int64, error) {
	connection, err := db.OpenConnection()
	if err != nil {
		return 0, nil
	}
	defer connection.Close()

	SQL := `UPDATE users SET name=$2, age=$3, gender=$4 WHERE id=$1`
	response, err := connection.Exec(SQL, id, user.Name, user.Age, user.Gender)
	if err != nil {
		return 0, nil
	}

	return response.RowsAffected()
}
