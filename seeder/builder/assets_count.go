package builder

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm"
)

func CreateAsset_count(db *gorm.DB, Plan_Code string, UserID string,
	Plan_Name string, TypeplanName string, MainBranchID string,
	Group string, Main_Category_ID string, Project_name string, Comment string,
	Plan_start string, Plan_end string, Status string, CategoryID string) error {
	status := &model.Asset_count{
		Asset_countID: uuid.New().String(),
		Plan_Code:     Plan_Code,
		Plan_Name:     &Plan_Name,
		UserID:        UserID,
		TypeplanName:  &TypeplanName,
		MainBranchID:  MainBranchID,
		Group:         Group,
		// Main_Category_ID: &Main_Category_ID,
		Project_name: &Project_name,
		Comment:      &Comment,
		Plan_start:   &Plan_start,
		Plan_end:     &Plan_end,
		Status:       &Status,
		// CategoryID:       &CategoryID,
	}

	if err := db.Create(status).Error; err != nil {
		fmt.Println("Error creating status:", err)
		return err
	}

	return nil
}
