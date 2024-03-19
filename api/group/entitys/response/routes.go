package GroupResult

import "time"

type GetAllResult struct {
	GroupID   string    `gorm:"type:uuid;primaryKey" json:"group_id"`
	GroupName string    `json:"group_name" validate:"required"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
	UserID    string    `json:"user_id"`
}
