package ServicesAsset_count

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetById(id string) (value *model.Asset_count, Error error) {
	db := database.DB
	inspection := model.Asset_count{Asset_countID: id}
	tx := db.Preload(clause.Associations).Find(&inspection)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &inspection, nil
}

func GetAllRecords() (result []*model.Asset_count, err error) {
	db := database.DB
	var records []*model.Asset_count
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}

// func CountLatestAsset_count() (count int, err error) {
// 	latestAsset_count, err := GetAllRecords()
// 	if err != nil {
// 		return 0, err
// 	}
// 	return len(latestAsset_count), nil
// }

func CountLatestAsset_count() (count int, err error) {
	latestAsset_count, err := GetAllRecords()
	if err != nil {
		return 0, err
	}

	for _, asset := range latestAsset_count {
		if asset.Status != nil && *asset.Status == "รอตรวจสอบ" {
			count++
		}
	}

	return count, nil
}

func CreateNewAsset_count(inspection model.Asset_count) (value *model.Asset_count, Error error) {
	db := database.DB
	inspection.Asset_countID = uuid.New().String()
	tx := db.Create(&inspection)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &inspection, nil
}

func UpdateAsset_count(id string, updatedType model.Asset_count) (value *model.Asset_count, Error error) {
	db := database.DB
	existingType := model.Asset_count{Asset_countID: id}
	tx := db.Model(&existingType).Where("asset_count_id = ?", id).Updates(updatedType)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &existingType, nil
}

func DeleteAsset_count(typeplan *model.Asset_count) error {
	db := database.DB
	tx := db.Delete(typeplan)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
