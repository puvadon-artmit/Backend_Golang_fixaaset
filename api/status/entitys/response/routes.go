package StatusResult

import "time"

type GetAllResult struct {
	StatusID   string    `gorm:"type:uuid;primaryKey" json:"status_id"`
	StatusName string    `json:"status_name" validate:"required"`
	Comment    string    `json:"comment"`
	CreatedAt  time.Time `json:"created_at"`
	UserID     string    `json:"user_id"`
}
