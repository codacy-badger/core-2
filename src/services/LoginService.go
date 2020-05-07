package services

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/iborg-ai/core/src/database"
	"github.com/iborg-ai/core/src/models"
)

// DoLogin hashes the password and queries the database
func DoLogin(email string, password string) models.HTTPErrorStatus {
	hash := md5.New()
	hash.Write([]byte(password))
	password = hex.EncodeToString(hash.Sum(nil))
	return database.DoesUserExistByEmailAndPassword(email, password)
}
