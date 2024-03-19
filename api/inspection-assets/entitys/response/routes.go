package response

import "time"

type GetAllResult struct {
	Asset_countID string     `json:"asset_count_id"`
	Plan_Code     string     `json:"plan_code"`
	Plan_Name     *string    `json:"plan_name"`
	TypeplanName  *string    `json:"typeplan_name"`
	BranchName    *string    `json:"name_branch"`
	Department    *string    `json:"department"`
	Property_type *string    `json:"property_type"`
	Project_name  *string    `json:"project_name"`
	Comment       *string    `json:"comment"`
	Plan_start    *string    `json:"plan_start"`
	Plan_end      *string    `json:"plan_end"`
	CreatedAt     *time.Time `json:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at"`
	UserID        string     `json:"user_id"`
	AssetsID      string     `json:"assets_id"`
}
