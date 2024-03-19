package RoleServices

import (
	"strings"

	"github.com/google/uuid"
	RoleResponse "github.com/puvadon-artmit/gofiber-template/api/role/entitys/response"
	"github.com/puvadon-artmit/gofiber-template/database"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm/clause"
)

func GetById(id string) (result *model.Role, Error error) {
	db := database.DB
	var role model.Role
	role.RoleID = id
	tx := db.Preload(clause.Associations).Find(&role)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &role, nil
}

func GetAll() (data *[]RoleResponse.GetAllResult, Error error) {
	db := database.DB
	var result []RoleResponse.GetAllResult
	tx := db.Table("roles").Select("role_id, role_name, role_display_name, role_description, created_at").Scan(&result)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &result, nil
}

func GenerateRoleCode() string {
	uuid := uuid.NewString()
	return strings.ToUpper(uuid[:34])
}

func CreateNewRole(role model.Role) (*model.Role, error) {
	db := database.DB
	role.RoleID = uuid.New().String()

	if role.RoleCode == "" {
		role.RoleCode = GenerateRoleCode()
	}

	err := db.Create(&role).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}
