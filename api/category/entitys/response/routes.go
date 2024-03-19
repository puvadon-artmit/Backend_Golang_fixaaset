package response

import "time"

type GetAllResult struct {
	CategoryID    string    `json:"category_id"`
	CategoryName  string    `json:"category_name"`
	CategoryImage string    `json:"category_image"`
	UserID        string    `json:"user_id"`
	CreatedAt     time.Time `json:"created_at"`
	ImageURL      string    `json:"image_url"`
}
