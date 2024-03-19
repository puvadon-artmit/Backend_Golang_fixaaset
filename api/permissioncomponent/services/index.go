package ServicesPermissionComponent

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetById(id string) (value *model.PermissionComponent, Error error) {
	db := database.DB
	group_story := model.PermissionComponent{PermissionComponentID: id}
	tx := db.Preload(clause.Associations).Find(&group_story)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &group_story, nil
}

func GetAllRecords() (result []*model.PermissionComponent, err error) {
	db := database.DB
	var records []*model.PermissionComponent
	tx := db.Preload(clause.Associations).Find(&records)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return records, nil
}

func CreateNewPermissionComponent(inspection model.PermissionComponent) (value *model.PermissionComponent, Error error) {
	db := database.DB
	inspection.PermissionComponentID = uuid.New().String()
	tx := db.Create(&inspection)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &inspection, nil
}

func UpdatePermissionComponent(id string, updatedType model.PermissionComponent) (value *model.PermissionComponent, Error error) {
	db := database.DB
	existingType := model.PermissionComponent{PermissionComponentID: id}
	tx := db.Model(&existingType).Where("permission_component_id = ?", id).Updates(updatedType)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &existingType, nil
}

func DeletePermissionComponent(group_story *model.PermissionComponent) error {
	db := database.DB
	tx := db.Delete(group_story)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
