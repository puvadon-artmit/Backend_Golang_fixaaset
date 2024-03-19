package response

import "time"

type GetAllResult struct {
	Typeplan_story     string    `json:"typeplan_story_id"`
	Typeplan_StoryName string    `json:"typeplan_story_name"`
	Typeplan_Username  string    `json:"typeplan_username"`
	Typeplan_Details   string    `json:"typeplan_details"`
	UserID             string    `json:"user_id"`
	TypeplanID         string    `json:"typeplan_id"`
	Comment            string    `json:"comment"`
	CreatedAt          time.Time `json:"created_at"`
}
