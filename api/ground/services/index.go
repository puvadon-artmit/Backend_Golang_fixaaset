package services

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetById(id string) (value *model.Ground, Error error) {
	db := database.DB
	ground := model.Ground{GroundID: id}
	tx := db.Preload(clause.Associations).Find(&ground)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &ground, nil
}

// func GetById(id string) (value *model.Category, Error error) {
// 	db := database.DB
// 	category := model.Category{CategoryID: id}
// 	tx := db.Preload(clause.Associations).Find(&category)
// 	if tx.Error != nil {
// 		return nil, tx.Error
// 	}
// 	return &category, nil
// }

func UpdateItemModelByID(id string, updatedItem model.Ground) error {
	db := database.DB

	existingItem := model.Ground{GroundID: id}
	tx := db.First(&existingItem)
	if tx.Error != nil {
		return tx.Error
	}
	tx = db.Model(&existingItem).Updates(updatedItem)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func GetAllRecords() (result []*model.Ground, err error) {
	db := database.DB
	var records []*model.Ground
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}

func CreateNewGround(ground model.Ground) (value *model.Ground, Error error) {
	db := database.DB
	ground.GroundID = uuid.New().String()
	tx := db.Create(&ground)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &ground, nil
}

func DeleteGroundPhoto(ground *model.Ground) error {
	db := database.DB
	tx := db.Delete(ground)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
