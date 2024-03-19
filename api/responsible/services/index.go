package ServicesResponsible

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetById(id string) (value *model.Responsible, Error error) {
	db := database.DB
	responsible := model.Responsible{ResponsibleID: id}
	tx := db.Preload(clause.Associations).Find(&responsible)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &responsible, nil
}

func GetAllRecords() (result []*model.Responsible, err error) {
	db := database.DB
	var records []*model.Responsible
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}

func CreateNewResponsible(responsible model.Responsible) (value *model.Responsible, Error error) {
	db := database.DB
	responsible.ResponsibleID = uuid.New().String()
	tx := db.Create(&responsible)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &responsible, nil
}
