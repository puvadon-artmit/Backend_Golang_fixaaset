package builder

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm"
)

func CreateResponsible(db *gorm.DB, ResponsibleName string, Comment string, UserID string, GroupID string, EmployeeCode string) error {
	return db.Create(&model.Responsible{
		ResponsibleID:   uuid.New().String(),
		ResponsibleName: &ResponsibleName,
		EmployeeCode:    &EmployeeCode,
		Comment:         &Comment,
		UserID:          UserID,
		GroupID:         GroupID,
	}).Error
}
