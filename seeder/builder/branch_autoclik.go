package builder

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm"
)

func CreateBranch_Autoclik(db *gorm.DB, Branch_Name string, Branch_Code string) error {
	status := &model.Branch_Autoclik{
		BranchAutoclik_ID: uuid.New().String(),
		Branch_Name:       Branch_Name,
		Branch_Code:       Branch_Code,
	}

	if err := db.Create(status).Error; err != nil {
		fmt.Println("Error creating status:", err)
		return err
	}

	return nil
}
