package builder

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm"
)

func CreateGround(db *gorm.DB, GroundName string, Comment string, UserID string, TypeID string, Location_code string,
	Building string, Floor string, Room string, GroundImage string) error {
	return db.Create(&model.Ground{
		GroundID:      uuid.New().String(),
		GroundName:    &GroundName,
		Comment:       &Comment,
		Location_code: &Location_code,
		Building:      &Building,
		Floor:         &Floor,
		Room:          &Room,
		GroundImage:   &GroundImage,
		UserID:        UserID,
	}).Error
}
