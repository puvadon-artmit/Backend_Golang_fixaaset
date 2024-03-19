package builder

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm"
)

func CreateType(db *gorm.DB, TypeName string, Comment string, UserID string, CategoryID string) error {
	return db.Create(&model.Type_things{
		TypeID:     uuid.New().String(),
		TypeName:   &TypeName,
		Comment:    &Comment,
		UserID:     UserID,
		CategoryID: CategoryID,
	}).Error
}
