package builder

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm"
)

func CreateRound_Count_Story(db *gorm.DB, Round_Name_Strory string, Round_Name string, UserID string, Round_Strory_Detail string, Round_CountID string) error {
	return db.Create(&model.Round_Count_Story{
		Round_Count_StoryID: uuid.New().String(),
		Round_Name_Strory:   &Round_Name_Strory,
		Round_Name:          &Round_Name,
		Round_Strory_Detail: &Round_Strory_Detail,
		UserID:              UserID,
		Round_CountID:       Round_CountID,
	}).Error
}
