package services

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetById(id string) (photosURL string, err error) {
	db := database.DB
	photos_check := model.Photos_check{Photos_checkID: id}
	tx := db.Select("photos").First(&photos_check)
	if tx.Error != nil {
		return "", tx.Error
	}
	return photos_check.Photos, nil
}

func GetByPhotoId(id string) (value *model.Photos_check, Error error) {
	db := database.DB
	getphoto := model.Photos_check{Photos_checkID: id}
	tx := db.Preload(clause.Associations).Find(&getphoto)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &getphoto, nil
}

func GetAll() (result []*model.Photos_check, err error) {
	db := database.DB
	var records []*model.Photos_check
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}

func CreateNewPhotos_check(photos_check model.Photos_check) (value *model.Photos_check, Error error) {
	db := database.DB
	photos_check.Photos_checkID = uuid.New().String()
	tx := db.Create(&photos_check)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &photos_check, nil
}

func DeletePhotos_check(photos_check *model.Photos_check) error {
	db := database.DB
	tx := db.Unscoped().Delete(photos_check)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
