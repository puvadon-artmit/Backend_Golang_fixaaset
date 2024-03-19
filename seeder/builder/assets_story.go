package builder

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm"
)

func CreateAssets_Story(db *gorm.DB, Assets_StoryName string, UserID string,
	Assets_Username string, AssetsID string, Assets_Details string) error {
	status := &model.Assets_Story{
		Assets_storyID:   uuid.New().String(),
		Assets_StoryName: &Assets_StoryName,
		Assets_Username:  &Assets_Username,
		UserID:           UserID,
		AssetsID:         AssetsID,
		Assets_Details:   &Assets_Details,
	}

	if err := db.Create(status).Error; err != nil {
		fmt.Println("Error creating status:", err)
		return err
	}

	return nil
}
