package PermissionResult

import "time"

type GetAllResult struct {
	PermissionID          string    `json:"permission_id"`
	PermissionName        string    `json:"permission_name"`
	PermissionDisplayName string    `json:"permission_display_name"`
	PermissionDescription string    `json:"permission_description"`
	CreatedAt             time.Time `json:"created_at"`
}
