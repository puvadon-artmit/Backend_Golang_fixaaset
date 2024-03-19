package builder

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm"
)

func CreateGroup_Story(db *gorm.DB, GroupStoryName string, UserID string,
	GroupStoryDetails string, GroupID string) error {

	status := &model.Group_Story{
		GroupStoryID:      uuid.New().String(),
		GroupStoryName:    GroupStoryName,
		GroupStoryDetails: &GroupStoryDetails,
		UserID:            UserID,
		GroupID:           GroupID,
	}

	if err := db.Create(status).Error; err != nil {
		fmt.Println("Error creating status:", err)
		return err
	}

	return nil
}
