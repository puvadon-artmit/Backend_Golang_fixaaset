package services

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetById(id string) (value *model.Item_model, Error error) {
	db := database.DB
	item_model := model.Item_model{ItemModelID: id}
	tx := db.Preload(clause.Associations).Find(&item_model)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &item_model, nil
}

func GetByItem_Photo(id string) (FrontpictureURL string, err error) {
	db := database.DB
	frontpicture := model.Item_model{ItemModelID: id}
	tx := db.Select("frontpicture").First(&frontpicture)
	if tx.Error != nil {
		return "", tx.Error
	}
	if frontpicture.Frontpicture != nil {
		return *frontpicture.Frontpicture, nil
	}
	return "", nil
}

func UpdateItemModelByID(id string, updatedItem model.Item_model) error {
	db := database.DB

	existingItem := model.Item_model{ItemModelID: id}
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

func GetAllRecords() (result []*model.Item_model, err error) {
	db := database.DB
	var records []*model.Item_model
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}

func CreateNewItem_model(item_model model.Item_model) (value *model.Item_model, Error error) {
	db := database.DB
	item_model.ItemModelID = uuid.New().String()
	tx := db.Create(&item_model)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &item_model, nil
}

func DeleteItem_model(item_model *model.Item_model) error {
	db := database.DB
	tx := db.Delete(item_model)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func DeleteUnusedImages() error {
	files, err := ioutil.ReadDir("./uploads/")
	if err != nil {
		return err
	}

	var itemModels []model.Item_model
	database.DB.Find(&itemModels)

	var categories []model.Category
	database.DB.Find(&categories)

	var grounds []model.Ground
	database.DB.Find(&grounds)

	var photos_check []model.Photos_check
	database.DB.Find(&photos_check)

	usedFilenames := make(map[string]bool)

	for _, itemModel := range itemModels {
		if itemModel.Frontpicture != nil {
			usedFilenames[*itemModel.Frontpicture] = true
		}
	}

	for _, category := range categories {
		if category.CategoryImage != "" {
			usedFilenames[category.CategoryImage] = true
		}
	}

	for _, ground := range grounds {
		if ground.GroundImage != nil {
			usedFilenames[*ground.GroundImage] = true
		}
	}

	for _, photo := range photos_check {
		if len(photo.Photos) > 0 {
			usedFilenames[photo.Photos] = true
		}
	}

	for _, file := range files {
		filename := file.Name()
		if _, exists := usedFilenames[filename]; !exists {
			err := os.Remove(filepath.Join("./uploads/", filename))
			if err != nil {
				return err
			}
		}
	}

	return nil
}
