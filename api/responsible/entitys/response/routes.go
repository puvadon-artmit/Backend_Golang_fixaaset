package response

import "time"

type GetAllResult struct {
	ResponsibleID   string    `json:"responsible_id"`
	ResponsibleName string    `json:"responsible_name"`
	EmployeeCode    string    `json:"employee_code"`
	Comment         string    `json:"comment"`
	UserID          string    `json:"user_id"`
	CreatedAt       time.Time `json:"created_at"`
	GroupID         string    `json:"group_id"`
}
