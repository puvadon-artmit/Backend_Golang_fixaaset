package ServicesAsset_count

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetById(id string) (value *model.Main_Branch_Story, Error error) {
	db := database.DB
	main_branch := model.Main_Branch_Story{MainBranchStoryID: id}
	tx := db.Preload(clause.Associations).Find(&main_branch)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &main_branch, nil
}

func GetAllRecords() (result []*model.Main_Branch_Story, err error) {
	db := database.DB
	var records []*model.Main_Branch_Story
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}

func CreateNewMain_Branch_Story(inspection model.Main_Branch_Story) (value *model.Main_Branch_Story, Error error) {
	db := database.DB
	inspection.MainBranchStoryID = uuid.New().String()
	tx := db.Create(&inspection)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &inspection, nil
}

func UpdateMain_Branch_Story(id string, updatedType model.Main_Branch_Story) (value *model.Main_Branch_Story, Error error) {
	db := database.DB
	existingType := model.Main_Branch_Story{MainBranchStoryID: id}
	tx := db.Model(&existingType).Where("main-branch-id = ?", id).Updates(updatedType)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &existingType, nil
}

func DeleteMain_Branch_Story(main_branch *model.Main_Branch_Story) error {
	db := database.DB
	tx := db.Delete(main_branch)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
