package services

import (
	"reflect"

	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetById(id string) (value *model.Asset_check, Error error) {
	db := database.DB
	asset_check := model.Asset_check{Asset_checkID: id}
	tx := db.Preload(clause.Associations).Find(&asset_check)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &asset_check, nil
}

// func GetAssetCheckByAssetCheckID(Round_CountID string) (assetCheck *model.Asset_check, err error) {
// 	db := database.DB
// 	var assetCheckModel model.Asset_check

// 	tx := db.Where("round_count_id = ?", Round_CountID).First(&assetCheckModel)
// 	if tx.Error != nil {
// 		return nil, tx.Error
// 	}

// 	return &assetCheckModel, nil
// }

// func GetAssetCheckByAssetCheckID(Round_CountID string) (assetChecks []*model.Asset_check, err error) {
// 	db := database.DB
// 	var assetCheckModels []*model.Asset_check

// 	tx := db.Where("round_count_id = ?", Round_CountID).Find(&assetCheckModels)
// 	if tx.Error != nil {
// 		return nil, tx.Error
// 	}

// 	return assetCheckModels, nil
// }

func GetAssetCheckByAssetCheckID(Round_CountID string) (assetChecks []*model.Asset_check, err error) {
	db := database.DB
	var assetCheckModels []*model.Asset_check

	tx := db.Where("round_count_id = ?", Round_CountID).Order("property_code desc").Find(&assetCheckModels)
	if tx.Error != nil {
		return nil, tx.Error
	}

	propertyCodes := make(map[string]bool)
	var filteredAssetChecks []*model.Asset_check

	for _, assetCheck := range assetCheckModels {
		if !propertyCodes[*assetCheck.Property_code] {
			propertyCodes[*assetCheck.Property_code] = true
			filteredAssetChecks = append(filteredAssetChecks, assetCheck)
		}
	}

	return filteredAssetChecks, nil
}

func UpdateAsset_checkByID(id string, updatedItem model.Asset_check) error {
	db := database.DB

	existingItem := model.Asset_check{Asset_checkID: id}
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

func GetAllRecords() (result []*model.Asset_check, err error) {
	db := database.DB
	var records []*model.Asset_check
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}

// func CreateNewAsset_check(asset_check model.Asset_check) (value *model.Asset_check, Error error) {
// 	db := database.DB
// 	asset_check.Asset_checkID = uuid.New().String()
// 	tx := db.Create(&asset_check)
// 	if tx.Error != nil {
// 		return nil, tx.Error
// 	}
// 	return &asset_check, nil
// }

func CreateNewAsset_check(asset_check model.Asset_check, photos_check []model.Photos_check) (value *model.Asset_check, Error error) {
	db := database.DB
	asset_check.Asset_checkID = uuid.New().String()

	tx := db.Create(&asset_check)
	if tx.Error != nil {
		return nil, tx.Error
	}

	fields := []string{"Additional_notes", "Property_code", "StatusAsset", "Round_CountID", "Photos_check", "UserID"}
	for _, field := range fields {
		if fieldValue := reflect.ValueOf(&asset_check).Elem().FieldByName(field); fieldValue.IsValid() && fieldValue.String() == "" {
			fieldValue.SetString("")
		}
	}

	for i := range photos_check {
		photoCheck := model.PhotoCheck{
			Asset_checkID:  asset_check.Asset_checkID,
			Photos_checkID: photos_check[i].Photos_checkID,
		}
		tx := db.Create(&photoCheck)
		if tx.Error != nil {
			return nil, tx.Error
		}
	}

	return &asset_check, nil
}

func GetByRound_CountID(Round_CountID string) ([]*model.Asset_check, error) {
	db := database.DB
	var round_count []*model.Asset_check
	tx := db.Preload(clause.Associations).Where("round_count_id = ?", Round_CountID).Find(&round_count)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return round_count, nil
}

// func CreateNewAsset_check(asset_check model.Asset_check) (*model.Asset_check, error) {
// 	db := database.DB

// 	assetCheckID := uuid.New()
// 	asset_check.Asset_checkID = assetCheckID.String()

// 	for i := range asset_check.Photos_check {
// 		photosCheckID := uuid.New()
// 		asset_check.Photos_check[i].Photos_checkID = photosCheckID.String()
// 	}

// 	currentTime := time.Now()
// 	asset_check.CreatedAt = &currentTime
// 	asset_check.UpdatedAt = &currentTime

// 	tx := db.Create(&asset_check)
// 	if tx.Error != nil {
// 		return nil, tx.Error
// 	}

// 	return &asset_check, nil
// }
