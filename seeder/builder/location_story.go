package builder

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm"
)

func CreateLocation_story(db *gorm.DB, Location_StoryName string, Location_Username string, Location_Details string, LocationID string) error {
	return db.Create(&model.Location_story{
		Location_storyID:   uuid.New().String(),
		Location_StoryName: &Location_StoryName,
		Location_Username:  &Location_Username,
		Location_Details:   &Location_Details,
		LocationID:         LocationID,
	}).Error
}
