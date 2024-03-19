package builder

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm"
)

func CreateMain_branch(db *gorm.DB, MainBranchName string) error {
	return db.Create(&model.Main_branch{
		MainBranchID:   uuid.New().String(),
		MainBranchName: &MainBranchName,
	}).Error
}
