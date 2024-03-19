package builder

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm"
)

func CreateItem_model(db *gorm.DB, ItemModelName string, Comment string, UserID string, TypeID string, ProductNumber int,
	Weight string, RequiredUnits int, Frontpicture string) error {
	return db.Create(&model.Item_model{
		ItemModelID:   uuid.New().String(),
		ItemModelName: &ItemModelName,
		Comment:       &Comment,
		ProductNumber: &ProductNumber,
		Weight:        &Weight,
		RequiredUnits: &RequiredUnits,
		Frontpicture:  &Frontpicture,
		UserID:        UserID,
		TypeID:        TypeID,
	}).Error
}
