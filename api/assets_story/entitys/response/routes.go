package Ground_StoryResult

import "time"

type GetAllResult struct {
	Assets_storyID   string     `gorm:"type:uuid;primaryKey" json:"assets_story_id"`
	Assets_StoryName *string    `json:"assetsstory_name"`
	Assets_Username  *string    `json:"assets_username"`
	Assets_Details   *string    `json:"assets_details"`
	CreatedAt        *time.Time `json:"created_at"`
	UpdatedAt        *time.Time `json:"updated_at"`
	UserID           string     `json:"user_id"`
	AssetsID         string     `json:"assets_id"`
}
