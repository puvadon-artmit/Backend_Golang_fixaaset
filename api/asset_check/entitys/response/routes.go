package response

import "time"

type GetAllResult struct {
	ItemModelName string     `json:"item_model_name"`
	Comment       string     `json:"comment"`
	ProductNumber int        `json:"product_number"`
	Weight        string     `json:"weight"`
	RequiredUnits int        `json:"required_units"`
	Frontpicture  string     `json:"frontpicture"`
	Rearpicture   string     `json:"rearpicture"`
	Otherpictures string     `json:"otherpictures"`
	CreatedAt     *time.Time `json:"created_at"`
	UserID        string     `json:"user_id"`
	TypeID        string     `json:"type_id"`
}
