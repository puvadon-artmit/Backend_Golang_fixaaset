package builder

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm"
)

func CreateGroup(db *gorm.DB, GroupName string, Comment string, UserID string) error {
	return db.Create(&model.Group{
		GroupID:   uuid.New().String(),
		GroupName: &GroupName,
		UserID:    UserID,
	}).Error
}
