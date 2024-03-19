package builder

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm"
)

func CreateTypeplan(db *gorm.DB, TypeplanName string, Comment string) error {
	return db.Create(&model.Typeplan{
		TypeplanID:   uuid.New().String(),
		TypeplanName: &TypeplanName,
		Comment:      &Comment,
	}).Error
}
