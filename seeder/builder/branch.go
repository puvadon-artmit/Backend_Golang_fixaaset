package builder

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm"
)

func CreateBranch(db *gorm.DB, GroupID string, ZipCode string, County string,
	Province string, Comment string, Building string, Address string, Town string,
	RoomNumber string, UserID string, MainBranchID string) error {
	return db.Create(&model.Branch{
		BranchID:     uuid.New().String(),
		GroupID:      GroupID,
		County:       &County,
		Province:     &Province,
		Comment:      &Comment,
		Building:     &Building,
		Address:      &Address,
		Town:         &Town,
		RoomNumber:   &RoomNumber,
		UserID:       UserID,
		MainBranchID: MainBranchID,
	}).Error
}
