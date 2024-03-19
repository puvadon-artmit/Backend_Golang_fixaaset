package BranchServices

import (
	"reflect"

	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetById(id string) (value *model.Branch, err error) {
	db := database.DB
	branch := model.Branch{BranchID: id}
	tx := db.Preload(clause.Associations).Find(&branch)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &branch, nil
}

func GetAll() (result []*model.Branch, err error) {
	db := database.DB
	var records []*model.Branch
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}

func CreateNewBranch(branch *model.Branch) (value *model.Branch, err error) {
	db := database.DB
	branch.BranchID = uuid.New().String()

	fields := []string{"ZipCode", "County", "Province", "Comment", "Building", "Address", "Town", "RoomNumber"}
	for _, field := range fields {
		if fieldValue := reflect.ValueOf(&branch).Elem().FieldByName(field); fieldValue.IsValid() && fieldValue.String() == "" {
			fieldValue.SetString("")
		}
	}

	tx := db.Create(branch)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return branch, nil
}
