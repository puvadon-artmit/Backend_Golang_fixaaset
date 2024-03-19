package builder

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm"
)

func CreateLocation(db *gorm.DB, Coordinates_x string, Coordinates_y string, AssetsID string, GroundID string) error {
	return db.Create(&model.Location{
		LocationID:    uuid.New().String(),
		Coordinates_x: Coordinates_x,
		Coordinates_y: Coordinates_y,
		AssetsID:      AssetsID,
		GroundID:      GroundID,
	}).Error
}
