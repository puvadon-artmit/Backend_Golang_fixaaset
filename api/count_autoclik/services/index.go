package ServicesCount_Autoclik

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetById(id string) (value *model.Count_Autoclik, Error error) {
	db := database.DB
	group_story := model.Count_Autoclik{Autoclik_countID: id}
	tx := db.Preload(clause.Associations).Find(&group_story)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &group_story, nil
}

func GetAllRecords() (result []*model.Count_Autoclik, err error) {
	db := database.DB
	var records []*model.Count_Autoclik
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}

func CreateNewCount_Autoclik(round_count model.Count_Autoclik) (value *model.Count_Autoclik, Error error) {
	db := database.DB
	round_count.Autoclik_countID = uuid.New().String()
	tx := db.Create(&round_count)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &round_count, nil
}

func UpdateCount_Autoclik(id string, updatedType model.Count_Autoclik) (value *model.Count_Autoclik, Error error) {
	db := database.DB
	existingType := model.Count_Autoclik{Autoclik_countID: id}
	tx := db.Model(&existingType).Where("autoclik_count_id = ?", id).Updates(updatedType)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &existingType, nil
}

func DeleteCount_Autoclik(group_story *model.Count_Autoclik) error {
	db := database.DB
	tx := db.Delete(group_story)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
