package response

import "time"

type GetAllResult struct {
	TypeID     string    `json:"type_id"`
	TypeName   string    `json:"type_name"`
	Comment    string    `json:"comment"`
	UserID     string    `json:"user_id"`
	CreatedAt  time.Time `json:"created_at"`
	CategoryID string    `json:"category_id"`
}
