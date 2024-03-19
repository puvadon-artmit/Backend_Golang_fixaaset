package builder

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm"
)

func CreateSignature(db *gorm.DB, SignatureName string, UserID string,
	Signature string, Asset_countID string) error {
	status := &model.Signature{
		SignatureID: uuid.New().String(),
		// SignatureName: &SignatureName,
		Signature:     &Signature,
		Asset_countID: Asset_countID,
		UserID:        UserID,
	}

	if err := db.Create(status).Error; err != nil {
		fmt.Println("Error creating status:", err)
		return err
	}

	return nil
}
