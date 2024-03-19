package ServicesRound_Count

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetById(id string) (value *model.Round_Count, Error error) {
	db := database.DB
	group_story := model.Round_Count{Round_CountID: id}
	tx := db.Preload(clause.Associations).Find(&group_story)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &group_story, nil
}

func GetAllRecords() (result []*model.Round_Count, err error) {
	db := database.DB
	var records []*model.Round_Count
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}

func GetDelete() (result []*model.Round_Count, err error) {
	db := database.DB
	var records []*model.Round_Count
	tx := db.Unscoped().Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}

func CreateNewRound_Count(round_count model.Round_Count) (value *model.Round_Count, Error error) {
	db := database.DB
	round_count.Round_CountID = uuid.New().String()
	tx := db.Create(&round_count)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &round_count, nil
}

func UpdateRound_Count(id string, updatedType model.Round_Count) (value *model.Round_Count, Error error) {
	db := database.DB
	existingType := model.Round_Count{Round_CountID: id}
	tx := db.Model(&existingType).Where("round_count_id = ?", id).Updates(updatedType)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &existingType, nil
}

func GetByAsset_countIDDB(Asset_countID string) ([]*model.Round_Count, error) {
	db := database.DB
	round_counts := []*model.Round_Count{}
	tx := db.Preload(clause.Associations).Where("asset_count_id = ?", Asset_countID).Find(&round_counts)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return round_counts, nil
}

func DeleteRound_Count(group_story *model.Round_Count) error {
	db := database.DB
	tx := db.Delete(group_story)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
