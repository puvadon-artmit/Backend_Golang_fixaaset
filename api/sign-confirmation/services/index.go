package services

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetById(id string) (value *model.Signature, Error error) {
	db := database.DB
	signature := model.Signature{SignatureID: id}
	tx := db.Preload(clause.Associations).Find(&signature)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &signature, nil
}

func UpdateSignatureByID(id string, updatedItem model.Signature) error {
	db := database.DB

	existingItem := model.Signature{SignatureID: id}
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

func GetAllRecords() (result []*model.Signature, err error) {
	db := database.DB
	var records []*model.Signature
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}

func CreateNewSignature(asset_check model.Signature) (value *model.Signature, Error error) {
	db := database.DB
	asset_check.SignatureID = uuid.New().String()
	tx := db.Create(&asset_check)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &asset_check, nil
}

func GetByAsset_countDB(Asset_countID string) ([]*model.Signature, error) {
	db := database.DB
	var asset_count []*model.Signature
	tx := db.Preload(clause.Associations).Where("asset_count_id = ?", Asset_countID).Find(&asset_count)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return asset_count, nil
}
