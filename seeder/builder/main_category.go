package builder

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm"
)

func CreateMain_Category(db *gorm.DB, Code_ID string, Main_name string) error {
	return db.Create(&model.Main_Category{
		Main_Category_ID: uuid.New().String(),
		Code_ID:          Code_ID,
		Main_name:        Main_name,
	}).Error
}
