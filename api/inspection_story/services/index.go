package ServicesInspection_story

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetById(id string) (value *model.Inspection_story, Error error) {
	db := database.DB
	inspection_story := model.Inspection_story{Inspection_storyID: id}
	tx := db.Preload(clause.Associations).Find(&inspection_story)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &inspection_story, nil
}

func GetAllRecords() (result []*model.Inspection_story, err error) {
	db := database.DB
	var records []*model.Inspection_story
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}

func CreateNewInspection_story(inspection_story model.Inspection_story) (value *model.Inspection_story, Error error) {
	db := database.DB
	inspection_story.Inspection_storyID = uuid.New().String()
	tx := db.Create(&inspection_story)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &inspection_story, nil
}

func UpdateInspection_story(id string, updatedType model.Inspection_story) (value *model.Inspection_story, Error error) {
	db := database.DB
	existingType := model.Inspection_story{Inspection_storyID: id}
	tx := db.Preload(clause.Associations).Find(&existingType)
	if tx.Error != nil {
		return nil, tx.Error
	}
	tx = db.Model(&existingType).Updates(updatedType)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &existingType, nil
}

func DeleteInspection_story(typeplan *model.Inspection_story) error {
	db := database.DB
	tx := db.Delete(typeplan)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
