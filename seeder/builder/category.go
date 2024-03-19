package builder

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm"
)

func CreateCategory(db *gorm.DB, CategoryName string, CategoryImage string, UserID string, Main_Category_ID string) error {
	return db.Create(&model.Category{
		CategoryID:       uuid.New().String(),
		CategoryName:     CategoryName,
		CategoryImage:    CategoryImage,
		UserID:           UserID,
		Main_Category_ID: Main_Category_ID,
	}).Error
}
