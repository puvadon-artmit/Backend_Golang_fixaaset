package response

import "time"

type GetAllResult struct {
	AssetsID             string    `json:"assets_id"`
	Model_name           *string   `json:"model_name"`
	Manufacturer         *string   `json:"manufacturer"`
	Serial_Code          *string   `json:"serial_code"`
	Type                 *string   `json:"type"`
	Model                *string   `json:"model"`
	Branch               *string   `json:"branch"`
	Username             *string   `json:"username"`
	Property_code        *string   `json:"property_code"`
	Status               *string   `json:"status"`
	Group_hardware       *string   `json:"group_hardware"`
	Group                *string   `json:"group"`
	User_hardware        *string   `json:"user_hardware"`
	Phone_number         *string   `json:"phone_number"`
	Posting_group        *string   `json:"posting_group"`
	Responsible_employee *string   `json:"responsible_employ"`
	Location_code        *string   `json:"location_code"`
	Comment1             *string   `json:"comment1"`
	Comment2             *string   `json:"comment2"`
	Comment3             *string   `json:"comment3"`
	ItemModelID          *string   `json:"item_model_id"`
	UserID               string    `json:"user_id"`
	CreatedAt            time.Time `json:"created_at"`
	CategoryID           string    `json:"category_id"`
}
