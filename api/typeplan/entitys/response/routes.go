package response

import "time"

type GetAllResult struct {
	TypeplanID   string    `json:"typeplan_id"`
	TypeplanName string    `json:"typeplan_name"`
	Comment      string    `json:"comment"`
	CreatedAt    time.Time `json:"created_at"`
}
