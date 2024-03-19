package builder

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm"
)

func CreateRound_Count(db *gorm.DB, Round string, Asset_countID string, Status string) error {
	return db.Create(&model.Round_Count{
		Round_CountID: uuid.New().String(),
		Round:         &Round,
		Status:        &Status,
		Asset_countID: Asset_countID,
	}).Error
}
