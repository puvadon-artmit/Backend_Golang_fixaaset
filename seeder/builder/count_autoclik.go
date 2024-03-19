package builder

import (
	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm"
)

func CreateCount_Autoclik(db *gorm.DB, Plan_Code string, Plan_Name string, UserID string,
	TypeplanName string, Project_name string, Comment string, Plan_start string, Plan_end string,
	Status string, Product_Group_Code string, Group string, BranchAutoclik_ID string,
) error {
	return db.Create(&model.Count_Autoclik{
		Autoclik_countID:   uuid.New().String(),
		Plan_Code:          Plan_Code,
		Plan_Name:          &Plan_Name,
		UserID:             UserID,
		TypeplanName:       &TypeplanName,
		Project_name:       &Project_name,
		Comment:            &Comment,
		Plan_start:         &Plan_start,
		Plan_end:           &Plan_end,
		Status:             &Status,
		Product_Group_Code: &Product_Group_Code,
		Group:              Group,
		BranchAutoclik_ID:  BranchAutoclik_ID,
	}).Error
}
