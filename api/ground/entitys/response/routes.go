package response

import "time"

type GetAllResult struct {
	GroundName    string     `json:"ground_name"`
	Location_code string     `json:"location_code"`
	Building      string     `json:"building"`
	Floor         string     `json:"floor"`
	Room          string     `json:"room"`
	GroundImage   string     `json:"ground_image"`
	Comment       string     `json:"comment"`
	CreatedAt     *time.Time `json:"created_at"`
	UserID        string     `json:"user_id"`
}
