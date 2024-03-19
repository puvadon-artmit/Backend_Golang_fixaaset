package response

import "time"

type GetAllResult struct {
	Location_storyID   string    `json:"location_story_id"`
	Location_StoryName *string   `json:"location_story_name"`
	Location_Username  *string   `json:"location_username"`
	Location_Details   *string   `json:"location_details"`
	CreatedAt          time.Time `json:"created_at"`
	LocationID         string    `json:"location_id"`
}
