package ServicesAsset_count_Main_Category

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetById(id string) (value *model.Asset_count_Main_Category, Error error) {
	db := database.DB
	types := model.Asset_count_Main_Category{Asset_count_Main_CategoryID: id}
	tx := db.Preload(clause.Associations).Find(&types)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &types, nil
}

func GetAllRecords() (result []*model.Asset_count_Main_Category, err error) {
	db := database.DB
	var records []*model.Asset_count_Main_Category
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}
func CreateNewcount_main_category(types model.Asset_count_Main_Category) (value *model.Asset_count_Main_Category, Error error) {
	db := database.DB
	types.Asset_count_Main_CategoryID = uuid.New().String()
	tx := db.Create(&types)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &types, nil
}

func Updatecount_main_category(id string, updatedType model.Asset_count_Main_Category) (value *model.Asset_count_Main_Category, Error error) {
	db := database.DB
	existingType := model.Asset_count_Main_Category{Asset_count_Main_CategoryID: id}
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

func Deletecount_main_category(typeplan *model.Asset_count_Main_Category) error {
	db := database.DB
	tx := db.Delete(typeplan)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func GetByAsset_countIDDB(Asset_countID string) ([]*model.Asset_count_Main_Category, error) {
	db := database.DB
	var asset_count []*model.Asset_count_Main_Category
	tx := db.Preload(clause.Associations).Where("asset_count_id = ?", Asset_countID).Find(&asset_count)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return asset_count, nil
}
