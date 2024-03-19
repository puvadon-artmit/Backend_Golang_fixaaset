package builder

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm"
)

func CreateBranch_Story(db *gorm.DB, BranchStoryName string, UserID string,
	BranchStoryDetails string, BranchID string) error {
	status := &model.Branch_Story{
		BranchStoryID:      uuid.New().String(),
		BranchStoryName:    BranchStoryName,
		BranchStoryDetails: BranchStoryDetails,
		UserID:             UserID,
		BranchID:           BranchID,
	}

	if err := db.Create(status).Error; err != nil {
		fmt.Println("Error creating status:", err)
		return err
	}

	return nil
}
