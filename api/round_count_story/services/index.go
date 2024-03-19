package ServicesRound_Count_Story

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetById(id string) (value *model.Round_Count_Story, Error error) {
	db := database.DB
	group_story := model.Round_Count_Story{Round_Count_StoryID: id}
	tx := db.Preload(clause.Associations).Find(&group_story)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &group_story, nil
}

func GetAllRecords() (result []*model.Round_Count_Story, err error) {
	db := database.DB
	var records []*model.Round_Count_Story
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}

func CreateNewRound_Count_Story(round_count model.Round_Count_Story) (value *model.Round_Count_Story, Error error) {
	db := database.DB
	round_count.Round_Count_StoryID = uuid.New().String()
	tx := db.Create(&round_count)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &round_count, nil
}

func UpdateRound_Count_Story(id string, updatedType model.Round_Count_Story) (value *model.Round_Count_Story, Error error) {
	db := database.DB
	existingType := model.Round_Count_Story{Round_Count_StoryID: id}
	tx := db.Model(&existingType).Where("round_count_id = ?", id).Updates(updatedType)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &existingType, nil
}

func DeleteRound_Count_Story(group_story *model.Round_Count_Story) error {
	db := database.DB
	tx := db.Delete(group_story)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
