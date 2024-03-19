package builder

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm"
)

func CreateResponsible_Story(db *gorm.DB, Responsible_StoryDetails string, UserID string,
	Responsible_StoryName string, ResponsibleID string) error {

	status := &model.Responsible_Story{
		ResponsibleStoryID:       uuid.New().String(),
		Responsible_StoryDetails: &Responsible_StoryDetails,
		Responsible_StoryName:    &Responsible_StoryName,
		UserID:                   UserID,
		ResponsibleID:            ResponsibleID,
	}

	if err := db.Create(status).Error; err != nil {
		fmt.Println("Error creating status:", err)
		return err
	}

	return nil
}
