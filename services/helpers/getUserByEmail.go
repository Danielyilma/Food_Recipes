package helpers

import (
	"database/sql"

	"github.com/Danielyilma/Food_Recipes/services/databases"
	"github.com/Danielyilma/Food_Recipes/services/models"
)

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	query := "select id, email, password from users where email=$1"
	row := databases.DBConnection.QueryRow(query, email)

	err := row.Scan(&user.Id, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err // Some other error
	}

	return &user, nil
}
