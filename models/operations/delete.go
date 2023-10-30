package operations

import (
	"api/db"
)

func DeleteUserByID(id int64) (int64, error) {
	connection, err := db.OpenConnection()
	if err != nil {
		return 0, nil
	}
	defer connection.Close()

	SQL := `DELETE FROM users WHERE id=$1`
	response, err := connection.Exec(SQL, id)
	if err != nil {
		return 0, nil
	}

	return response.RowsAffected()
}
