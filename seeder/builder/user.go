package builder

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm"
)

func CreateUsers(db *gorm.DB, Username string, Firstname string, Lastname string, Password string, RoleID string) error {
	return db.Create(&model.User{UserID: uuid.New().String(), Username: Username, Firstname: Firstname, Lastname: Lastname, Password: Password}).Error
}
