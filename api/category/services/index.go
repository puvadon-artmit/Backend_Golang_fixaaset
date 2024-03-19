package services

import (
	"reflect"

	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetById(id string) (value *model.Category, Error error) {
	db := database.DB
	category := model.Category{CategoryID: id}
	tx := db.Preload(clause.Associations).Find(&category)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &category, nil
}

// func GetAll() (value *[]response.GetAllResult, Error error) {
// 	db := database.DB
// 	var result []response.GetAllResult
// 	// tx := db.Table("categories").Select("category_id, category_name, category_image, user_id, created_at").Scan(&result)
// 	tx := db.Preload(clause.Associations).Find(&result)
// 	if tx.Error != nil {
// 		return nil, tx.Error
// 	}
// 	return &result, nil
// }

func GetAll() (result []*model.Category, err error) {
	db := database.DB
	var records []*model.Category
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}

func CountLatestCategory() (count int, err error) {
	latestCategory, err := GetAll()
	if err != nil {
		return 0, err
	}
	return len(latestCategory), nil
}

func CreateNewCategory(category model.Category) (value *model.Category, Error error) {
	db := database.DB
	category.CategoryID = uuid.New().String()

	fields := []string{"CategoryName", "CategoryImage", "Main_Category_ID", "UserID"}
	for _, field := range fields {
		if fieldValue := reflect.ValueOf(&category).Elem().FieldByName(field); fieldValue.IsValid() && fieldValue.String() == "" {
			fieldValue.SetString("")
		}
	}

	tx := db.Create(&category)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &category, nil
}

func GetByMain_Category_IDDB(Main_Category_ID string) ([]*model.Category, error) {
	db := database.DB
	var autocliks []*model.Category
	tx := db.Preload(clause.Associations).Where("main_category_id = ?", Main_Category_ID).Find(&autocliks)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return autocliks, nil
}

// func GetByCategory_Photos(id string) (FrontpictureURL string, err error) {
// 	db := database.DB
// 	category_image := model.Category{CategoryID: id}
// 	tx := db.Select("category_image").First(&category_image)
// 	if tx.Error != nil {
// 		return "", tx.Error
// 	}
// 	if category_image.CategoryImage != nil {
// 		return *category_image.Frontpicture, nil
// 	}
// 	return "", nil
// }

// CategoryImage

func GetByCategory_Photo(id string) (CategoryImageURL string, err error) {
	db := database.DB
	categoryimage := model.Category{CategoryID: id}
	tx := db.Select("photos").First(&categoryimage)
	if tx.Error != nil {
		return "", tx.Error
	}
	return categoryimage.CategoryImage, nil
}

func GetByCategory_Photos(id string) (value *model.Category, Error error) {
	db := database.DB
	getphoto := model.Category{CategoryID: id}
	tx := db.Preload(clause.Associations).Find(&getphoto)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &getphoto, nil
}

func DeleteCategory(category *model.Category) error {
	db := database.DB
	tx := db.Delete(category)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
