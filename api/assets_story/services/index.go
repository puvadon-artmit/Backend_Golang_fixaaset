package Assets_StoryServices

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetById(id string) (value *model.Assets_Story, err error) {
	db := database.DB
	story := model.Assets_Story{Assets_storyID: id}
	tx := db.Preload(clause.Associations).Find(&story)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &story, nil
}

func GetAllRecords() (result []*model.Assets_Story, err error) {
	db := database.DB
	var records []*model.Assets_Story
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}

func CreateNewStory(status *model.Assets_Story) (value *model.Assets_Story, err error) {
	db := database.DB
	status.Assets_storyID = uuid.New().String()
	tx := db.Create(status)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return status, nil
}
