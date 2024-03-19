package StoryServices

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetById(id string) (value *model.Story, err error) {
	db := database.DB
	story := model.Story{StoryID: id}
	tx := db.Preload(clause.Associations).Find(&story)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &story, nil
}

func GetAllRecords() (result []*model.Story, err error) {
	db := database.DB
	var records []*model.Story
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}

func CreateNewStory(status *model.Story) (value *model.Story, err error) {
	db := database.DB
	status.StoryID = uuid.New().String()
	tx := db.Create(status)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return status, nil
}
