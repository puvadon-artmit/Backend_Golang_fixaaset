package builder

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm"
)

func CreateAsset_Round(db *gorm.DB, AssetsID string, UserID string,
	Photos string, Asset_countID string, Round_CountID string) error {
	status := &model.Asset_Round{
		Asset_RoundID: uuid.New().String(),
		AssetsID:      AssetsID,
		Asset_countID: Asset_countID,
		Round_CountID: Round_CountID,
	}

	if err := db.Create(status).Error; err != nil {
		fmt.Println("Error creating status:", err)
		return err
	}

	return nil
}
