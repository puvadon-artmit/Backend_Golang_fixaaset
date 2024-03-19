package ServicesType

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetById(id string) (value *model.Typeplan_story, Error error) {
	db := database.DB
	types := model.Typeplan_story{Typeplan_storyID: id}
	tx := db.Preload(clause.Associations).Find(&types)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &types, nil
}

func GetAllRecords() (result []*model.Typeplan_story, err error) {
	db := database.DB
	var records []*model.Typeplan_story
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}
func CreateNewTypeplan_story(types model.Typeplan_story) (value *model.Typeplan_story, Error error) {
	db := database.DB
	types.Typeplan_storyID = uuid.New().String()
	tx := db.Create(&types)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &types, nil
}

func UpdateTypeplan_story(id string, updatedType model.Typeplan_story) (value *model.Typeplan_story, Error error) {
	db := database.DB
	existingType := model.Typeplan_story{Typeplan_storyID: id}
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

func DeleteTypeplan_story(typeplan *model.Typeplan_story) error {
	db := database.DB
	tx := db.Delete(typeplan)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
