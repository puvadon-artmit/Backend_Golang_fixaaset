package ServicesResponsible_Story

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetById(id string) (value *model.Responsible_Story, Error error) {
	db := database.DB
	group_story := model.Responsible_Story{ResponsibleStoryID: id}
	tx := db.Preload(clause.Associations).Find(&group_story)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &group_story, nil
}

func GetAllRecords() (result []*model.Responsible_Story, err error) {
	db := database.DB
	var records []*model.Responsible_Story
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}

func CreateNewResponsible_Story(inspection model.Responsible_Story) (value *model.Responsible_Story, Error error) {
	db := database.DB
	inspection.ResponsibleStoryID = uuid.New().String()
	tx := db.Create(&inspection)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &inspection, nil
}

func UpdateResponsible_Story(id string, updatedType model.Responsible_Story) (value *model.Responsible_Story, Error error) {
	db := database.DB
	existingType := model.Responsible_Story{ResponsibleStoryID: id}
	tx := db.Model(&existingType).Where("responsible_story_id = ?", id).Updates(updatedType)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &existingType, nil
}

func DeleteResponsible_Story(group_story *model.Responsible_Story) error {
	db := database.DB
	tx := db.Delete(group_story)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
