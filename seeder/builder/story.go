package builder

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm"
)

func CreateStory(db *gorm.DB, StoryName string, Comment string, UserID string,
	Username string, ItemModelID string, Details string) error {
	status := &model.Story{
		StoryID:     uuid.New().String(),
		StoryName:   &StoryName,
		Username:    &Username,
		UserID:      UserID,
		ItemModelID: ItemModelID,
		Details:     &Details,
	}

	if err := db.Create(status).Error; err != nil {
		fmt.Println("Error creating status:", err)
		return err
	}

	return nil
}
