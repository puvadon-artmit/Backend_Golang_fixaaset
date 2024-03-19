package builder

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm"
)

func CreateStatus(db *gorm.DB, StatusName string, Comment string, UserID string) error {
	status := &model.Status{
		StatusID:   uuid.New().String(),
		StatusName: &StatusName,
		UserID:     UserID,
	}

	if err := db.Create(status).Error; err != nil {
		fmt.Println("Error creating status:", err)
		return err
	}

	return nil
}
