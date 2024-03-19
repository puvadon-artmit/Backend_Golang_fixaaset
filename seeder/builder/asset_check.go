package builder

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm"
)

func CreateAsset_check(db *gorm.DB, Additional_notes string, UserID string, Round_CountID uuid.UUID,
	StatusAsset string, Property_code string) error {
	status := &model.Asset_check{
		Asset_checkID:    uuid.New().String(),
		Additional_notes: &Additional_notes,
		Round_CountID:    Round_CountID,
		StatusAsset:      &StatusAsset,
		Property_code:    &Property_code,
	}

	if err := db.Create(status).Error; err != nil {
		fmt.Println("Error creating status:", err)
		return err
	}

	return nil
}
