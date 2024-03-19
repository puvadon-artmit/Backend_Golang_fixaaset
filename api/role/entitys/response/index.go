package RoleResponse

import "time"

type GetAllResult struct {
	RoleID          string    `json:"role_id"`
	RoleName        string    `json:"role_name"`
	RoleDisplayName string    `json:"role_display_name"`
	RoleDescription string    `json:"role_description"`
	CreatedAt       time.Time `json:"created_at"`
}
