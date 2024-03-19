package ServicesLocation_story

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetById(id string) (value *model.Location_story, Error error) {
	db := database.DB
	types := model.Location_story{Location_storyID: id}
	tx := db.Preload(clause.Associations).Find(&types)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &types, nil
}

func GetAllRecords() (result []*model.Location_story, err error) {
	db := database.DB
	var records []*model.Location_story
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}

func CreateNewLocation(location_story model.Location_story) (value *model.Location_story, Error error) {
	db := database.DB
	location_story.Location_storyID = uuid.New().String()
	tx := db.Create(&location_story)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &location_story, nil
}

func UpdateLocation_story(id string, location_story model.Location_story) (value *model.Location_story, Error error) {
	db := database.DB
	existingType := model.Location_story{Location_storyID: id}
	tx := db.Preload(clause.Associations).Find(&existingType)
	if tx.Error != nil {
		return nil, tx.Error
	}
	tx = db.Model(&existingType).Updates(location_story)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &existingType, nil
}

func DeleteLocation_story(typeplan *model.Location_story) error {
	db := database.DB
	tx := db.Delete(typeplan)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
