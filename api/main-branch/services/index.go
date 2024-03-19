package ServicesAsset_count

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetById(id string) (value *model.Main_branch, Error error) {
	db := database.DB
	main_branch := model.Main_branch{MainBranchID: id}
	tx := db.Preload(clause.Associations).Find(&main_branch)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &main_branch, nil
}

func GetAllRecords() (result []*model.Main_branch, err error) {
	db := database.DB
	var records []*model.Main_branch
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}

func CreateNewMain_branch(inspection model.Main_branch) (value *model.Main_branch, Error error) {
	db := database.DB
	inspection.MainBranchID = uuid.New().String()
	tx := db.Create(&inspection)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &inspection, nil
}

func UpdateMain_branch(id string, updatedType model.Main_branch) (value *model.Main_branch, Error error) {
	db := database.DB
	existingType := model.Main_branch{MainBranchID: id}
	tx := db.Model(&existingType).Where("main-branch-id = ?", id).Updates(updatedType)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &existingType, nil
}

func DeleteMain_branch(main_branch *model.Main_branch) error {
	db := database.DB
	tx := db.Delete(main_branch)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
