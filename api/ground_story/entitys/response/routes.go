package Ground_StoryResult

import "time"

type GetAllResult struct {
	Ground_storyID   string    `gorm:"type:uuid;primaryKey" json:"ground_story_id"`
	Ground_StoryName string    `json:"groundstory_name" validate:"required"`
	Ground_Username  string    `json:"ground_username"`
	Ground_Details   string    `json:"ground_details"`
	CreatedAt        time.Time `json:"created_at"`
	UserID           string    `json:"user_id"`
	GroundID         string    `json:"ground_id"`
}
