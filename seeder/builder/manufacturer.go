package builder

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm"
)

func CreateManufacturer(db *gorm.DB, ManufacturerName string, Comment string, UserID string) error {
	return db.Create(&model.Manufacturer{
		ManufacturerID:   uuid.New().String(),
		ManufacturerName: &ManufacturerName,
		Comment:          &Comment,
		UserID:           UserID,
	}).Error
}
