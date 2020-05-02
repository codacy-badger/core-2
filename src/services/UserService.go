package services

import (
	"crypto/md5"
	"encoding/hex"
	"time"

	"github.com/pravinba9495/iborg-core/src/database"
	"github.com/pravinba9495/iborg-core/src/models"
)

// SaveUser saves the user
func SaveUser(user models.User) int {
	hasher := md5.New()
	hasher.Write([]byte(user.Password))
	user.Password = hex.EncodeToString(hasher.Sum(nil))
	user.CreatedAt = time.Now().Unix()
	user.UpdatedAt = user.CreatedAt
	user.ID = ""
	return database.SaveUser(user)
}
