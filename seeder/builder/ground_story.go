package builder

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm"
)

func CreateGround_Story(db *gorm.DB, Ground_StoryName string, Comment string, UserID string,
	Ground_Username string, GroundID string, Ground_Details string) error {
	status := &model.Ground_Story{
		Ground_storyID:   uuid.New().String(),
		Ground_StoryName: &Ground_StoryName,
		Ground_Username:  &Ground_Username,
		UserID:           UserID,
		GroundID:         GroundID,
		Ground_Details:   &Ground_Details,
	}

	if err := db.Create(status).Error; err != nil {
		fmt.Println("Error creating status:", err)
		return err
	}

	return nil
}
