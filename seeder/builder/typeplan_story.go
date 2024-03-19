package builder

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm"
)

func CreateTypeplan_story(db *gorm.DB, Typeplan_StoryName string, Typeplan_Username string, UserID string, Typeplan_Details string, TypeplanID string) error {
	return db.Create(&model.Typeplan_story{
		Typeplan_storyID:   uuid.New().String(),
		Typeplan_StoryName: &Typeplan_StoryName,
		Typeplan_Username:  &Typeplan_Username,
		Typeplan_Details:   &Typeplan_Details,
		UserID:             UserID,
		TypeplanID:         TypeplanID,
	}).Error
}
