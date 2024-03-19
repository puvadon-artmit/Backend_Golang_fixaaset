package StoryResult

import "time"

type GetAllResult struct {
	StoryID     string    `gorm:"type:uuid;primaryKey" json:"story_id"`
	StoryName   string    `json:"story_name" validate:"required"`
	Username    string    `json:"username"`
	Details     string    `json:"details"`
	CreatedAt   time.Time `json:"created_at"`
	UserID      string    `json:"user_id"`
	ItemModelID string    `json:"item_model_id"`
}
