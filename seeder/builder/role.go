package builder

import (
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm"
)

func CreateRole(db *gorm.DB, RoleID string, RoleName string, RoleDisplayName string, RoleDescription string, RoleCode string) error {
	return db.Create(&model.Role{RoleID: RoleID, RoleName: RoleName, RoleDisplayName: RoleDisplayName, RoleDescription: RoleDescription, RoleCode: RoleCode}).Error
}
