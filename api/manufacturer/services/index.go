package ServicesManufacturer

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetById(id string) (value *model.Manufacturer, Error error) {
	db := database.DB
	manufacturer := model.Manufacturer{ManufacturerID: id}
	tx := db.Preload(clause.Associations).Find(&manufacturer)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &manufacturer, nil
}

func GetAllRecords() (result []*model.Manufacturer, err error) {
	db := database.DB
	var records []*model.Manufacturer
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}

func CreateNewManufacturer(manufacturer model.Manufacturer) (value *model.Manufacturer, Error error) {
	db := database.DB
	manufacturer.ManufacturerID = uuid.New().String()
	tx := db.Create(&manufacturer)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &manufacturer, nil
}
