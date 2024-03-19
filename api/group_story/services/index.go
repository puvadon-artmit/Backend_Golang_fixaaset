package ServicesGroup_Story

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetById(id string) (value *model.Group_Story, Error error) {
	db := database.DB
	group_story := model.Group_Story{GroupStoryID: id}
	tx := db.Preload(clause.Associations).Find(&group_story)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &group_story, nil
}

func GetAllRecords() (result []*model.Group_Story, err error) {
	db := database.DB
	var records []*model.Group_Story
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}

func CreateNewGroup_Story(inspection model.Group_Story) (value *model.Group_Story, Error error) {
	db := database.DB
	inspection.GroupStoryID = uuid.New().String()
	tx := db.Create(&inspection)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &inspection, nil
}

func UpdateGroup_Story(id string, updatedType model.Group_Story) (value *model.Group_Story, Error error) {
	db := database.DB
	existingType := model.Group_Story{GroupStoryID: id}
	tx := db.Model(&existingType).Where("group_story_id = ?", id).Updates(updatedType)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &existingType, nil
}

func DeleteGroup_Story(group_story *model.Group_Story) error {
	db := database.DB
	tx := db.Delete(group_story)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
