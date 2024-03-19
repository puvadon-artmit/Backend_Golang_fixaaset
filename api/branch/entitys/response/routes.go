package BranchResult

import "time"

type GetAllResult struct {
	BranchID   string    `json:"id_branch"`
	BranchName string    `json:"name_branch"`
	Subnet     string    `json:"subnet"`
	ZipCode    string    `json:"zip_code"`
	County     string    `json:"county"`
	Province   string    `json:"province"`
	Comment    string    `json:"comment"`
	Building   string    `json:"building"`
	Address    string    `json:"address"`
	Town       string    `json:"town"`
	RoomNumber string    `json:"room_number"`
	UserID     string    `json:"user_id"`
	CreatedAt  time.Time `json:"created_at"`
}
