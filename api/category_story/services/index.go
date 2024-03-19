package ServicesCategory_Story

import (
	"reflect"

	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetById(id string) (value *model.Category_Story, Error error) {
	db := database.DB
	category := model.Category_Story{Category_Story_ID: id}
	tx := db.Preload(clause.Associations).Find(&category)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &category, nil
}

func GetAllRecords() (result []*model.Category_Story, err error) {
	db := database.DB
	var getcategory []*model.Category_Story
	tx := db.Preload(clause.Associations).Find(&getcategory)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return getcategory, nil
}

func CreateNewCategory_Story(types model.Category_Story) (value *model.Category_Story, Error error) {
	db := database.DB
	types.Category_Story_ID = uuid.New().String()

	fields := []string{"Category_Story_Dtail", "Category_Story_Name", "CategoryID", "UserID"}
	for _, field := range fields {
		if fieldValue := reflect.ValueOf(&types).Elem().FieldByName(field); fieldValue.IsValid() && fieldValue.String() == "" {
			fieldValue.SetString("")
		}
	}

	tx := db.Create(&types)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &types, nil
}

func UpdateCategory_Story(id string, updatedType model.Category_Story) (value *model.Category_Story, Error error) {
	db := database.DB
	Updategetcategory := model.Category_Story{Category_Story_ID: id}
	tx := db.Preload(clause.Associations).Find(&Updategetcategory)
	if tx.Error != nil {
		return nil, tx.Error
	}
	tx = db.Model(&Updategetcategory).Updates(updatedType)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &Updategetcategory, nil
}

func DeleteCategory_Story(typeplan *model.Category_Story) error {
	db := database.DB
	tx := db.Delete(typeplan)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
