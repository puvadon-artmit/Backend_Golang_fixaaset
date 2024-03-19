package builder

import (
	"time"

	"github.com/google/uuid"
	"github.com/puvadon-artmit/gofiber-template/model"
	"gorm.io/gorm"
)

func CreateAssets(db *gorm.DB, Model_name string, Comment1 string, UserID string, CategoryID string,
	Manufacturer string, Serial_Code string, Type string, Model string, Branch string, Username string,
	Property_code string, Status string, Group_hardware string, Group string, User_hardware string,
	Phone_number string, Posting_group string, Responsible_employee string, GroundID string,
	Comment2 string, Comment3 string, ItemModelID string, ResponsibleID string, Latest_time time.Time) error {
	return db.Create(&model.Assets{
		AssetsID:       uuid.New().String(),
		Model_name:     Model_name,
		Comment1:       Comment1,
		Manufacturer:   Manufacturer,
		Serial_Code:    Serial_Code,
		Latest_time:    Latest_time,
		UserID:         UserID,
		CategoryID:     CategoryID,
		ItemModelID:    ItemModelID,
		Type:           Type,
		Model:          Model,
		Branch:         Branch,
		Username:       Username,
		Property_code:  Property_code,
		Status:         Status,
		Group_hardware: Group_hardware,
		Group:          Group,
		User_hardware:  User_hardware,
		Phone_number:   Phone_number,
		Posting_group:  Posting_group,
		ResponsibleID:  ResponsibleID,
		GroundID:       GroundID,
		Comment2:       Comment2,
		Comment3:       Comment3,
	}).Error
}
