package database

import (
	"github.com/pravinba9495/iborg-core/src/models"
	"net/http"
)

// SaveUser saves a user in the database
func SaveUser(user models.User) models.HttpErrorStatus {
	db := DBConn
	if db.NewRecord(user) {
		errors := db.Create(&user).GetErrors()
		if len(errors) == 0 {
			return models.HttpErrorStatus { Status: 201, Error: http.StatusText(201)}
		} else {
			return models.HttpErrorStatus { Status: 400, Error: errors[0].Error()}
		}
	} else {
		return models.HttpErrorStatus { Status: 409, Error: "Primary field present in body"}
	}
}
