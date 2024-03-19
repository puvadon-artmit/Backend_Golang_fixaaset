package StatusServices

import (
	"github.com/google/uuid"
	StatusResult "github.com/puvadon-artmit/gofiber-template/api/status/entitys/response"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetById(id string) (value *model.Status, err error) {
	db := database.DB
	status := model.Status{StatusID: id}
	tx := db.Preload(clause.Associations).Find(&status)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &status, nil
}

func GetAll() (value *[]StatusResult.GetAllResult, err error) {
	db := database.DB
	var result []StatusResult.GetAllResult
	tx := db.Table("statuses").Select("status_id, status_name, comment, created_at, user_id").Scan(&result)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &result, nil
}

func CreateNewStatus(status *model.Status) (value *model.Status, err error) {
	db := database.DB
	status.StatusID = uuid.New().String()
	tx := db.Create(status)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return status, nil
}
