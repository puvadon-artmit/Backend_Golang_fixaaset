package GroupServices

import (
	"github.com/google/uuid"
	GroupResult "github.com/puvadon-artmit/gofiber-template/api/group/entitys/response"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetById(id string) (value *model.Group, err error) {
	db := database.DB
	group := model.Group{GroupID: id}
	tx := db.Preload(clause.Associations).Find(&group)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &group, nil
}

func GetAll() (value *[]GroupResult.GetAllResult, err error) {
	db := database.DB
	var result []GroupResult.GetAllResult
	tx := db.Table("groups").Select("group_id, group_name, comment, created_at, user_id").Scan(&result)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &result, nil
}

func CreateNewGroup(group *model.Group) (value *model.Group, err error) {
	db := database.DB
	group.GroupID = uuid.New().String()
	tx := db.Create(group)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return group, nil
}
