package PermissionServices

import (
	"github.com/google/uuid"
	PermissionResult "github.com/puvadon-artmit/gofiber-template/api/permission/entitys/response"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetById(id string) (value *model.Permission, Error error) {
	db := database.DB
	permission := model.Permission{PermissionID: id}
	tx := db.Preload(clause.Associations).Find(&permission)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &permission, nil
}

func GetAll() (value *[]PermissionResult.GetAllResult, Error error) {
	db := database.DB
	var result []PermissionResult.GetAllResult
	tx := db.Table("permissions").Select("permission_id, permission_name, permission_display_name, permission_description, created_at").Scan(&result)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &result, nil
}

func CreateNewPermission(permission model.Permission) (value *model.Permission, Error error) {
	db := database.DB
	permission.PermissionID = uuid.New().String()
	tx := db.Create(&permission)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &permission, nil
}
