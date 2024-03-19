package ServicesBranch_Story

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetById(id string) (value *model.Branch_Story, Error error) {
	db := database.DB
	inspection_story := model.Branch_Story{BranchStoryID: id}
	tx := db.Preload(clause.Associations).Find(&inspection_story)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &inspection_story, nil
}

func GetAllRecords() (result []*model.Branch_Story, err error) {
	db := database.DB
	var records []*model.Branch_Story
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}

func CreateNewBranch_Story(inspection_story model.Branch_Story) (value *model.Branch_Story, Error error) {
	db := database.DB
	inspection_story.BranchStoryID = uuid.New().String()
	tx := db.Create(&inspection_story)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &inspection_story, nil
}

func UpdateBranch_Story(id string, updatedType model.Branch_Story) (value *model.Branch_Story, Error error) {
	db := database.DB
	existingType := model.Branch_Story{BranchStoryID: id}
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

func DeleteBranch_Story(typeplan *model.Branch_Story) error {
	db := database.DB
	tx := db.Delete(typeplan)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
