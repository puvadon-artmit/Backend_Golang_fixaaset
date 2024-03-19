package Ground_StoryServices

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetById(id string) (value *model.Ground_Story, err error) {
	db := database.DB
	story := model.Ground_Story{Ground_storyID: id}
	tx := db.Preload(clause.Associations).Find(&story)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &story, nil
}

func GetAllRecords() (result []*model.Ground_Story, err error) {
	db := database.DB
	var records []*model.Ground_Story
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}

func CreateNewStory(status *model.Ground_Story) (value *model.Ground_Story, err error) {
	db := database.DB
	status.Ground_storyID = uuid.New().String()
	tx := db.Create(status)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return status, nil
}
