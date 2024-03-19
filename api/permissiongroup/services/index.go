package ServicesPermissionGroup

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetById(id string) (value *model.PermissionGroup, Error error) {
	db := database.DB
	permissiongroup := model.PermissionGroup{PermissionGroupID: id}
	tx := db.Preload(clause.Associations).Find(&permissiongroup)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &permissiongroup, nil
}

func GetAllRecords() (result []*model.PermissionGroup, err error) {
	db := database.DB
	var records []*model.PermissionGroup
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}

func CreateNewPermissionGroup(permissiongroup model.PermissionGroup) (value *model.PermissionGroup, Error error) {
	db := database.DB
	permissiongroup.PermissionGroupID = uuid.New().String()
	tx := db.Create(&permissiongroup)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &permissiongroup, nil
}

func UpdatePermissionGroup(id string, updatedType model.PermissionGroup) (value *model.PermissionGroup, Error error) {
	db := database.DB
	existingType := model.PermissionGroup{PermissionGroupID: id}
	tx := db.Model(&existingType).Where("permission_group_id = ?", id).Updates(updatedType)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &existingType, nil
}

func DeletePermissionGroup(permissiongroup *model.PermissionGroup) error {
	db := database.DB
	tx := db.Delete(permissiongroup)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
