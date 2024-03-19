package ServicesAsset_Round

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetById(id string) (value *model.Asset_Round, Error error) {
	db := database.DB
	group_story := model.Asset_Round{Asset_RoundID: id}
	tx := db.Preload(clause.Associations).Find(&group_story)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &group_story, nil
}

func GetAllRecords() (result []*model.Asset_Round, err error) {
	db := database.DB
	var records []*model.Asset_Round
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}

func CreateNewAsset_Round(round_count model.Asset_Round) (value *model.Asset_Round, Error error) {
	db := database.DB
	round_count.Asset_RoundID = uuid.New().String()
	tx := db.Create(&round_count)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &round_count, nil
}

func UpdateAsset_Round(id string, updatedType model.Asset_Round) (value *model.Asset_Round, Error error) {
	db := database.DB
	existingType := model.Asset_Round{Asset_RoundID: id}
	tx := db.Model(&existingType).Where("asset_round_id = ?", id).Updates(updatedType)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &existingType, nil
}

func DeleteAsset_Round(group_story *model.Asset_Round) error {
	db := database.DB
	tx := db.Delete(group_story)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
