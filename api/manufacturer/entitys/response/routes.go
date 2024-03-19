package response

import "time"

type GetAllResult struct {
	ManufacturerID   string    `json:"manufacturer_id"`
	ManufacturerName string    `json:"manufacturer_name"`
	Comment          string    `json:"comment"`
	UserID           string    `json:"user_id"`
	CreatedAt        time.Time `json:"created_at"`
	TypeID           string    `json:"type_id"`
}
