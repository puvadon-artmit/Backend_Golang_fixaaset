package builder

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm"
)

func CreateCategory_Story(db *gorm.DB, Category_Story_Name string, Category_Story_Dtail string, UserID string, CategoryID string) error {
	return db.Create(&model.Category_Story{
		Category_Story_ID:    uuid.New().String(),
		Category_Story_Name:  Category_Story_Name,
		Category_Story_Dtail: Category_Story_Dtail,
		UserID:               UserID,
		CategoryID:           CategoryID,
	}).Error
}
