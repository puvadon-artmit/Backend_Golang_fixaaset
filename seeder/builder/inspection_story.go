package builder

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm"
)

func CreateInspection_story(db *gorm.DB, Inspection_name string, Inspection_Username string,
	Inspection_Details string, Asset_countID string) error {

	status := &model.Inspection_story{
		Inspection_storyID:  uuid.New().String(),
		Inspection_name:     &Inspection_name,
		Inspection_Username: &Inspection_Username,
		Inspection_Details:  &Inspection_Details,
		Asset_countID:       Asset_countID,
	}

	if err := db.Create(status).Error; err != nil {
		fmt.Println("Error creating status:", err)
		return err
	}

	return nil
}
